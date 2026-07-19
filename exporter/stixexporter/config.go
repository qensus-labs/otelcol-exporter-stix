package stixexporter

import "go.opentelemetry.io/collector/component"

// Config defines the STIX exporter configuration.
type Config struct {
	component.Config `mapstructure:",squash"`

	// Output controls where STIX bundles are written.
	Output string `mapstructure:"output"`
}
