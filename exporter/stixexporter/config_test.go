package stixexporter

import (
	"testing"

	"go.opentelemetry.io/collector/confmap"

	"github.com/qensus-labs/otelcol-exporter-stix/exporter/stixexporter/taxii"
)

func TestUnmarshalTAXIIConfig(t *testing.T) {

	conf := confmap.NewFromStringMap(map[string]any{
		"mode": "taxii",
		"taxii": map[string]any{
			"api_root":      "https://taxii.example.com/root",
			"collection_id": "collection-id",
			"username":      "admin",
			"password":      "secret",
			"api_key":       "token",
		},
	})

	cfg := createDefaultConfig().(*Config)

	if err := conf.Unmarshal(cfg); err != nil {
		t.Fatal(err)
	}

	want := taxii.Config{
		APIRoot:      "https://taxii.example.com/root",
		CollectionID: "collection-id",
		Username:     "admin",
		Password:     "secret",
		APIKey:       "token",
	}

	if cfg.TAXII != want {
		t.Fatalf("got %+v, want %+v", cfg.TAXII, want)
	}
}

func TestValidateStdout(t *testing.T) {

	cfg := &Config{
		Mode: "stdout",
	}

	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestValidateFile(t *testing.T) {

	cfg := &Config{
		Mode:   "file",
		Output: "bundle.json",
	}

	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestValidateFileMissingOutput(t *testing.T) {

	cfg := &Config{
		Mode: "file",
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal(
			"expected validation error",
		)
	}
}

func TestValidateTAXII(t *testing.T) {

	cfg := &Config{
		Mode: "taxii",

		TAXII: taxii.Config{
			APIRoot: "https://taxii.example.com/taxii2/root",

			CollectionID: "collection-id",
		},
	}

	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestValidateMissingAPIRoot(t *testing.T) {

	cfg := &Config{
		Mode: "taxii",

		TAXII: taxii.Config{
			CollectionID: "collection-id",
		},
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal(
			"expected validation error",
		)
	}
}

func TestValidateMissingCollectionID(t *testing.T) {

	cfg := &Config{
		Mode: "taxii",

		TAXII: taxii.Config{
			APIRoot: "https://taxii.example.com/taxii2/root",
		},
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal(
			"expected validation error",
		)
	}
}

func TestValidateUnknownMode(t *testing.T) {

	cfg := &Config{
		Mode: "banana",
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal(
			"expected validation error",
		)
	}
}
