package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
)

type schema struct {
	Ref                  string            `json:"$ref"`
	Type                 string            `json:"type"`
	Description          string            `json:"description"`
	Format               string            `json:"format"`
	Enum                 []string          `json:"enum"`
	Properties           map[string]schema `json:"properties"`
	Required             []string          `json:"required"`
	AnyOf                []schema          `json:"anyOf"`
	OneOf                json.RawMessage   `json:"oneOf"`
	AllOf                json.RawMessage   `json:"allOf"`
	Items                *schema           `json:"items"`
	AdditionalProperties json.RawMessage   `json:"additionalProperties"`
	Deprecated           bool              `json:"deprecated"`
}

type source struct {
	APIVersion   string `json:"api_version"`
	SourceCommit string `json:"source_commit"`
	SHA256       string `json:"sha256"`
}

type generator struct{ schemas map[string]schema }

func main() {
	if err := generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generate() error {
	data, err := os.ReadFile("api/openapi.json")
	if err != nil {
		return err
	}
	lockData, err := os.ReadFile("api/source.json")
	if err != nil {
		return err
	}
	var lock source
	if err := json.Unmarshal(lockData, &lock); err != nil {
		return err
	}
	sum := sha256.Sum256(data)
	if hex.EncodeToString(sum[:]) != lock.SHA256 {
		return fmt.Errorf("OpenAPI SHA-256 differs from api/source.json")
	}
	var spec struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
		Components struct {
			Schemas map[string]schema `json:"schemas"`
		} `json:"components"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		return err
	}
	if spec.Info.Version != lock.APIVersion {
		return fmt.Errorf("OpenAPI version differs from source lock")
	}
	g := generator{schemas: spec.Components.Schemas}
	var out bytes.Buffer
	fmt.Fprintf(&out, "// Code generated from API %s, SHA-256 %s; DO NOT EDIT.\n\npackage adanos\n\nimport \"encoding/json\"\n\n", lock.APIVersion, lock.SHA256)
	names := sortedKeys(g.schemas)
	for _, name := range names {
		s := g.schemas[name]
		if s.Type != "object" || len(s.Properties) == 0 {
			return fmt.Errorf("%s: expected a named object schema", name)
		}
		if s.OneOf != nil || s.AllOf != nil || len(s.AnyOf) > 0 {
			return fmt.Errorf("%s: unsupported object composition", name)
		}
		if goName(name) != name {
			return fmt.Errorf("%s: unsupported model name", name)
		}
		if s.Description != "" {
			comment(&out, name, s.Description)
		}
		fmt.Fprintf(&out, "type %s struct {\n", name)
		required := map[string]bool{}
		for _, key := range s.Required {
			required[key] = true
		}
		used := map[string]string{}
		for _, key := range sortedKeys(s.Properties) {
			field := goName(key)
			if prior, ok := used[field]; ok {
				return fmt.Errorf("%s: %s and %s collide as %s", name, prior, key, field)
			}
			used[field] = key
			property := s.Properties[key]
			typ, err := g.goType(property, !required[key], name+"."+key)
			if err != nil {
				return err
			}
			description := property.Description
			if property.Format != "" {
				description += " Format: " + property.Format + "."
			}
			if len(property.Enum) > 0 {
				description += " Values: " + strings.Join(property.Enum, ", ") + "."
			}
			if description != "" {
				comment(&out, field, description)
			}
			if property.Deprecated {
				fmt.Fprintln(&out, "// Deprecated: Retained for API compatibility.")
			}
			tag := key
			if !required[key] {
				tag += ",omitempty"
			}
			fmt.Fprintf(&out, "%s %s `json:%q`\n", field, typ, tag)
		}
		fmt.Fprintln(&out, "}")
		fmt.Fprintln(&out)
	}
	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("format generated models: %w", err)
	}
	return os.WriteFile("models.gen.go", formatted, 0644)
}

func (g generator) goType(s schema, optional bool, path string) (string, error) {
	if s.OneOf != nil || s.AllOf != nil {
		return "", fmt.Errorf("%s: unsupported schema composition", path)
	}
	if len(s.AnyOf) > 0 {
		var nonnull []schema
		nullable := false
		for _, variant := range s.AnyOf {
			if variant.Type == "null" {
				nullable = true
			} else {
				nonnull = append(nonnull, variant)
			}
		}
		if len(nonnull) > 1 {
			switch path {
			case "InvalidPeriodErrorDetail.value", "InvalidSearchParametersError.detail", "ValidationError.loc[]":
				return "json.RawMessage", nil
			}
			return "", fmt.Errorf("%s: unsupported mixed union", path)
		}
		if len(nonnull) != 1 {
			return "", fmt.Errorf("%s: expected exactly one non-null variant", path)
		}
		return g.goType(nonnull[0], optional || nullable, path)
	}
	var typ string
	if s.Ref != "" {
		name := strings.TrimPrefix(s.Ref, "#/components/schemas/")
		if name == s.Ref {
			return "", fmt.Errorf("%s: external schema ref %s", path, s.Ref)
		}
		if _, ok := g.schemas[name]; !ok {
			return "", fmt.Errorf("%s: unresolved ref %s", path, name)
		}
		typ = name
	} else {
		switch s.Type {
		case "string":
			typ = "string"
		case "integer":
			typ = "int64"
		case "number":
			typ = "float64"
		case "boolean":
			typ = "bool"
		case "array":
			if s.Items == nil {
				return "", fmt.Errorf("%s: array has no items", path)
			}
			item, err := g.goType(*s.Items, false, path+"[]")
			if err != nil {
				return "", err
			}
			return "[]" + item, nil
		case "object":
			if len(s.Properties) > 0 {
				return "", fmt.Errorf("%s: unsupported inline object", path)
			}
			if bytes.Equal(s.AdditionalProperties, []byte("true")) {
				return "map[string]json.RawMessage", nil
			}
			if len(s.AdditionalProperties) == 0 || bytes.Equal(s.AdditionalProperties, []byte("false")) {
				return "", fmt.Errorf("%s: object has no typed additional properties", path)
			}
			var value schema
			if err := json.Unmarshal(s.AdditionalProperties, &value); err != nil {
				return "", err
			}
			item, err := g.goType(value, false, path+"{}")
			if err != nil {
				return "", err
			}
			return "map[string]" + item, nil
		default:
			return "", fmt.Errorf("%s: unsupported schema type %q", path, s.Type)
		}
	}
	if optional {
		typ = "*" + typ
	}
	return typ, nil
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func goName(name string) string {
	parts := strings.FieldsFunc(name, func(r rune) bool { return r == '_' || r == '-' })
	initialisms := map[string]string{"id": "ID", "url": "URL", "api": "API", "utc": "UTC", "http": "HTTP"}
	for i, part := range parts {
		if upper, ok := initialisms[part]; ok {
			parts[i] = upper
		} else if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func comment(out *bytes.Buffer, name, text string) {
	text = strings.Join(strings.Fields(text), " ")
	if text != "" {
		fmt.Fprintf(out, "// %s: %s\n", name, text)
	}
}
