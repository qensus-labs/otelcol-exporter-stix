package stixexporter

import (
	"go.opentelemetry.io/collector/component"

	"github.com/qensus-labs/otelcol-exporter-stix/exporter/stixexporter/taxii"
)

// Config defines the STIX exporter configuration.
type Config struct {
	component.Config `mapstructure:",squash"`

	// Sender mode:
	//
	// stdout
	// file
	// taxii
	Mode string `mapstructure:"mode"`

	// Output controls where STIX bundles are written.
	//
	// Used when mode=file.
	Output string `mapstructure:"output"`

	// TAXII 2.1 configuration.
	//
	// Used when mode=taxii.
	TAXII taxii.Config `mapstructure:"taxii"`
}
