package adanos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// APIError describes a non-successful HTTP response. Payload is the original
// response body, including partial bytes when reading the body fails.
type APIError struct {
	StatusCode int
	Detail     string
	Payload    json.RawMessage
	Headers    http.Header
	cause      error
}

func (e *APIError) Error() string { return fmt.Sprintf("adanos: HTTP %d: %s", e.StatusCode, e.Detail) }
func (e *APIError) Unwrap() error { return e.cause }

func execute[T any](c *Client, ctx context.Context, method, path string, query url.Values, body any) (T, error) {
	var out T
	if err := c.request(ctx, method, path, query, body, &out); err != nil {
		var zero T
		return zero, err
	}
	return out, nil
}

func (c *Client) request(ctx context.Context, method, path string, query url.Values, body, out any) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("adanos: request URL: %w", err)
	}
	u.RawQuery = query.Encode()
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("adanos: encode request: %w", err)
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return fmt.Errorf("adanos: create request: %w", err)
	}
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("adanos: request failed: %w", err)
	}
	defer res.Body.Close()
	data, readErr := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return &APIError{StatusCode: res.StatusCode, Detail: errorDetail(data, res.StatusCode), Payload: json.RawMessage(data), Headers: res.Header.Clone(), cause: readErr}
	}
	if readErr != nil {
		return fmt.Errorf("adanos: read response: %w", readErr)
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("adanos: decode HTTP %d response: %w", res.StatusCode, err)
	}
	return nil
}

func errorDetail(data []byte, status int) string {
	var payload any
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if !json.Valid(data) || decoder.Decode(&payload) != nil {
		return http.StatusText(status)
	}
	if m, ok := payload.(map[string]any); ok {
		if detail, ok := m["detail"]; ok {
			payload = detail
		}
	}
	return detailText(payload, status)
}

func detailText(value any, status int) string {
	switch v := value.(type) {
	case string:
		if v != "" {
			return v
		}
	case map[string]any:
		for _, key := range []string{"message", "error", "detail"} {
			if item, ok := v[key]; ok {
				if text := detailText(item, status); text != "" && text != http.StatusText(status) {
					return text
				}
			}
		}
	case []any:
		var messages []string
		for _, item := range v {
			if issue, ok := item.(map[string]any); ok {
				if msg, ok := issue["msg"].(string); ok {
					var locations []string
					if loc, ok := issue["loc"].([]any); ok {
						for _, part := range loc {
							locations = append(locations, fmt.Sprint(part))
						}
					}
					if len(locations) > 0 {
						msg = strings.Join(locations, ".") + ": " + msg
					}
					messages = append(messages, msg)
				}
			}
		}
		if len(messages) > 0 {
			return strings.Join(messages, "; ")
		}
	}
	if value != nil {
		if raw, err := json.Marshal(value); err == nil {
			return string(raw)
		}
	}
	return http.StatusText(status)
}

func segment(value string) string {
	if value == "." {
		return "%2E"
	}
	if value == ".." {
		return "%2E%2E"
	}
	return url.PathEscape(value)
}
