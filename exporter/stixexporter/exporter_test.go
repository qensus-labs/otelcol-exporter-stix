package stixexporter

import (
	"bytes"
	"context"
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"
)

type bufferSender struct {
	buffer bytes.Buffer
}

func (s *bufferSender) Send(
	ctx context.Context,
	data []byte,
) error {

	_, err := s.buffer.Write(data)

	return err
}

func (s *bufferSender) Close() error {

	return nil
}

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

	sender := &bufferSender{}

	exporter := &logsExporter{
		sender: sender,
	}

	err := exporter.consumeLogs(
		context.Background(),
		logs,
	)

	if err != nil {
		t.Fatal(err)
	}

	if sender.buffer.Len() == 0 {
		t.Fatal(
			"expected STIX bundle output",
		)
	}

	output := sender.buffer.String()

	if !bytes.Contains(
		sender.buffer.Bytes(),
		[]byte(`"type": "bundle"`),
	) {

		t.Fatalf(
			"expected STIX bundle JSON, got: %s",
			output,
		)
	}
}
