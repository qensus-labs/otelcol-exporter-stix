package stixexporter

import (
	"bytes"
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

	var buffer bytes.Buffer

	exporter := &logsExporter{
		writer: &buffer,
	}

	err := exporter.consumeLogs(
		context.Background(),
		logs,
	)

	if err != nil {
		t.Fatal(err)
	}

	if buffer.Len() == 0 {
		t.Fatal("expected STIX bundle output")
	}

	output := buffer.String()

	if !bytes.Contains(
		buffer.Bytes(),
		[]byte(`"type": "bundle"`),
	) {

		t.Fatalf(
			"expected STIX bundle JSON, got: %s",
			output,
		)
	}
}
