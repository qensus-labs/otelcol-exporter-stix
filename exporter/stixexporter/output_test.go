package stixexporter

import (
	"os"
	"testing"
)

func TestCreateWriterStdout(t *testing.T) {

	writer, err := createWriter(
		"stdout",
	)

	if err != nil {
		t.Fatal(err)
	}

	if writer != os.Stdout {
		t.Fatal(
			"expected stdout writer",
		)
	}
}

func TestCreateWriterDefault(t *testing.T) {

	writer, err := createWriter(
		"",
	)

	if err != nil {
		t.Fatal(err)
	}

	if writer != os.Stdout {
		t.Fatal(
			"expected default writer to be stdout",
		)
	}
}

func TestCreateWriterFile(t *testing.T) {

	path := "test-stix-output.json"

	defer os.Remove(path)

	writer, err := createWriter(
		path,
	)

	if err != nil {
		t.Fatal(err)
	}

	if writer == nil {
		t.Fatal(
			"expected file writer",
		)
	}

	_, ok := writer.(*os.File)

	if !ok {
		t.Fatal(
			"expected *os.File writer",
		)
	}

	// Cleanup open file handle.
	writer.(*os.File).Close()
}
