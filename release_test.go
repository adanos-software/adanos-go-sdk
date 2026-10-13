package adanos_test

import (
	"encoding/json"
	"os"
	"testing"

	adanos "github.com/adanos-software/adanos-go-sdk"
)

func TestReleaseMetadata(t *testing.T) {
	data, err := os.ReadFile("api/source.json")
	if err != nil {
		t.Fatal(err)
	}
	var source struct {
		APIVersion string `json:"api_version"`
	}
	if err := json.Unmarshal(data, &source); err != nil {
		t.Fatal(err)
	}
	if adanos.APIVersion != source.APIVersion {
		t.Fatalf("SDK API version %s differs from frozen contract %s", adanos.APIVersion, source.APIVersion)
	}
	if tag := os.Getenv("ADANOS_RELEASE_TAG"); tag != "" && tag != "v"+adanos.Version {
		t.Fatalf("release tag %s differs from SDK version %s", tag, adanos.Version)
	}
}
