package stixexporter

import (
	"context"
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"
)

func TestConsumeLogsCreatesSTIXBundle(t *testing.T) {

	logs := plog.NewLogs()

	resourceLogs := logs.ResourceLogs().AppendEmpty()

	scopeLogs := resourceLogs.ScopeLogs().AppendEmpty()

	record := scopeLogs.LogRecords().AppendEmpty()

	record.Attributes().PutStr(
		"client.address",
		"10.0.0.5",
	)

	record.Attributes().PutStr(
		"url.full",
		"https://example.com",
	)

	record.Attributes().PutStr(
		"process.name",
		"nginx",
	)

	record.Attributes().PutInt(
		"process.pid",
		1234,
	)

	exporter := &logsExporter{
		output: "stdout",
	}

	err := exporter.consumeLogs(
		context.Background(),
		logs,
	)

	if err != nil {
		t.Fatal(err)
	}
}
