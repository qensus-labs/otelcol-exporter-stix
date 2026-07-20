package stixexporter

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"

	"github.com/qensus-labs/otelcol-exporter-stix/exporter/stixexporter/taxii"
)

// Config defines the STIX exporter configuration.
type Config struct {
	component.Config `mapstructure:",squash"`

	// Sender mode.
	//
	// stdout
	// file
	// taxii
	Mode string `mapstructure:"mode"`

	// Used when mode=file.
	Output string `mapstructure:"output"`

	// Used when mode=taxii.
	TAXII taxii.Config `mapstructure:"taxii"`
}

// Validate validates the exporter configuration.
func (cfg *Config) Validate() error {

	switch cfg.Mode {

	case "", "stdout":

		return nil

	case "file":

		if cfg.Output == "" {
			return errors.New(
				"output is required when mode=file",
			)
		}

		return nil

	case "taxii":

		if cfg.TAXII.APIRoot == "" {
			return fmt.Errorf(
				"mode=taxii requires taxii.api_root",
			)
		}

		if cfg.TAXII.CollectionID == "" {
			return errors.New(
				"taxii.collection_id is required",
			)
		}

		return nil

	default:

		return fmt.Errorf(
			"unsupported sender mode %q",
			cfg.Mode,
		)

	}
}
