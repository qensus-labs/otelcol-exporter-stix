package stixexporter

import (
	"context"
	"testing"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/pdata/plog"
)

type testSender struct {
	sent bool
	data []byte
}

func (s *testSender) Send(
	ctx context.Context,
	data []byte,
) error {

	s.sent = true
	s.data = data

	return nil
}

func (s *testSender) Close() error {
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

	sender := &testSender{}

	exporter := &logsExporter{
		sender: sender,
		logger: zap.NewNop(),
	}

	err := exporter.consumeLogs(
		context.Background(),
		logs,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !sender.sent {
		t.Fatal(
			"expected sender to receive STIX bundle",
		)
	}

	if len(sender.data) == 0 {
		t.Fatal(
			"expected STIX bundle data",
		)
	}
}
