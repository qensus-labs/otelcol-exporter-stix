package stixexporter

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
)

type logsExporter struct {
}

func newLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {

	exp := &logsExporter{}

	return exporterhelper.NewLogs(
		ctx,
		set,
		cfg,
		exp.consumeLogs,
		exporterhelper.WithCapabilities(
			consumer.Capabilities{
				MutatesData: false,
			},
		),
	)
}

func (e *logsExporter) consumeLogs(
	ctx context.Context,
	logs plog.Logs,
) error {

	// Commit 4:
	// Convert plog.Logs into STIX objects.

	return nil
}
