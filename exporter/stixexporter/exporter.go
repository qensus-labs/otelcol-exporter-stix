package stixexporter

import (
	"context"
	"fmt"

	"github.com/qensus-labs/go-stix/stix"
	stixotel "github.com/qensus-labs/go-stix/stix/mapping/otel"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
)

type logsExporter struct {
	output string
}

func newLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {

	exp := &logsExporter{
		output: cfg.Output,
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

	switch e.output {

	case "", "stdout":

		fmt.Println(string(data))

	default:

		return fmt.Errorf(
			"unsupported STIX output: %s",
			e.output,
		)
	}

	return nil
}
