package stixexporter

import (
	"context"
	"os"
	"testing"
)

func TestLogsExporterShutdownClosesFile(t *testing.T) {

	path := "test-output.json"

	file, err := os.Create(path)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(path)

	exp := &logsExporter{
		sender: &fileSender{
			file: file,
		},
	}

	err = exp.Shutdown(
		context.Background(),
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = file.Write(
		[]byte("test"),
	)

	if err == nil {
		t.Fatal(
			"expected closed file",
		)
	}
}
