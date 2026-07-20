package stixexporter

import (
	"context"
	"fmt"

	"github.com/qensus-labs/go-stix/stix"
	stixotel "github.com/qensus-labs/go-stix/stix/mapping/otel"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
)

type logsExporter struct {
	sender Sender

	logger *zap.Logger
}

func newLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {

	sender, err := createSender(cfg)

	if err != nil {
		return nil, err
	}

	logger := set.Logger

	if logger == nil {
		logger = zap.NewNop()
	}

	exp := &logsExporter{
		sender: sender,
		logger: logger,
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

func (e *logsExporter) log() *zap.Logger {

	if e.logger == nil {
		return zap.NewNop()
	}

	return e.logger
}

func (e *logsExporter) Start(
	ctx context.Context,
	host component.Host,
) error {

	e.log().Info(
		"Starting STIX exporter",
	)

	return nil
}

func (e *logsExporter) Shutdown(
	ctx context.Context,
) error {

	e.log().Info(
		"Stopping STIX exporter",
	)

	if e.sender != nil {

		return e.sender.Close()

	}

	return nil
}

func (e *logsExporter) consumeLogs(
	ctx context.Context,
	logs plog.Logs,
) error {

	resourceLogs := logs.ResourceLogs()

	e.log().Debug(
		"Converting OpenTelemetry logs to STIX",
		zap.Int(
			"resource_logs",
			resourceLogs.Len(),
		),
	)

	builder := stix.NewBuilder()

	scopeLogCount := 0
	logRecordCount := 0

	for i := 0; i < resourceLogs.Len(); i++ {

		scopeLogs := resourceLogs.At(i).ScopeLogs()

		scopeLogCount += scopeLogs.Len()

		for j := 0; j < scopeLogs.Len(); j++ {

			logRecords := scopeLogs.At(j).LogRecords()

			logRecordCount += logRecords.Len()

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

					e.log().Error(
						"Failed to map log record to STIX",
						zap.Error(err),
					)

					return fmt.Errorf(
						"map log record to STIX: %w",
						err,
					)
				}
			}
		}
	}

	e.log().Debug(
		"Processed OpenTelemetry logs",
		zap.Int(
			"scope_logs",
			scopeLogCount,
		),
		zap.Int(
			"log_records",
			logRecordCount,
		),
	)

	data, err := builder.JSON()

	if err != nil {

		e.log().Error(
			"Failed to generate STIX bundle",
			zap.Error(err),
		)

		return fmt.Errorf(
			"generate STIX bundle: %w",
			err,
		)
	}

	e.log().Debug(
		"Generated STIX bundle",
		zap.Int(
			"bundle_bytes",
			len(data),
		),
	)

	e.log().Debug(
		"Sending STIX bundle",
	)

	err = e.sender.Send(
		ctx,
		data,
	)

	if err != nil {

		e.log().Error(
			"Failed to export STIX bundle",
			zap.Error(err),
		)

		return fmt.Errorf(
			"send STIX bundle: %w",
			err,
		)
	}

	e.log().Info(
		"Successfully exported STIX bundle",
		zap.Int(
			"log_records",
			logRecordCount,
		),
		zap.Int(
			"bundle_bytes",
			len(data),
		),
	)

	return nil
}
