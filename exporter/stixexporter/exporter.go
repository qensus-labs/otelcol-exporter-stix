package stixexporter

import (
	"context"
	"io"

	"github.com/qensus-labs/go-stix/stix"
	stixotel "github.com/qensus-labs/go-stix/stix/mapping/otel"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
)

type logsExporter struct {
	writer io.Writer
}

func newLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {

	writer, err := createWriter(
		cfg.Output,
	)

	if err != nil {
		return nil, err
	}

	exp := &logsExporter{
		writer: writer,
	}

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

	builder := stix.NewBuilder()

	resourceLogs := logs.ResourceLogs()

	for i := 0; i < resourceLogs.Len(); i++ {

		scopeLogs := resourceLogs.At(i).ScopeLogs()

		for j := 0; j < scopeLogs.Len(); j++ {

			logRecords := scopeLogs.At(j).LogRecords()

			for k := 0; k < logRecords.Len(); k++ {

				record := logRecords.At(k)

				obs := stixotel.FromLogRecord(
					record,
				)

				err := stixotel.MapObservation(
					builder,
					obs,
				)

				if err != nil {
					return err
				}
			}
		}
	}

	data, err := builder.JSON()

	if err != nil {
		return err
	}

	_, err = e.writer.Write(data)

	return err
}
