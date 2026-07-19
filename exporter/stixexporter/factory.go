package stixexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

var typeStr = component.MustNewType("stix")

// NewFactory creates a new STIX exporter factory.
func NewFactory() exporter.Factory {

	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(
			createLogsExporter,
			component.StabilityLevelAlpha,
		),
	)
}

func createDefaultConfig() component.Config {

	return &Config{}
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {

	return newLogsExporter(
		ctx,
		set,
		cfg.(*Config),
	)
}
