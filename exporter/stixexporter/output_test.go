package stixexporter

import (
	"os"
	"testing"
)

func TestCreateWriterStdout(t *testing.T) {

	writer, closer, err := createWriter(
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

	if closer != nil {
		t.Fatal(
			"stdout should not have a closer",
		)
	}
}

func TestCreateWriterDefault(t *testing.T) {

	writer, closer, err := createWriter(
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

	if closer != nil {
		t.Fatal(
			"stdout should not have a closer",
		)
	}
}

func TestCreateWriterFile(t *testing.T) {

	path := "test-stix-output.json"

	defer os.Remove(path)

	writer, closer, err := createWriter(
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

	if closer == nil {
		t.Fatal(
			"expected file closer",
		)
	}

	file, ok := writer.(*os.File)

	if !ok {
		t.Fatal(
			"expected *os.File writer",
		)
	}

	err = closer.Close()

	if err != nil {
		t.Fatal(err)
	}

	// Verify the file is closed.
	_, err = file.Write(
		[]byte("test"),
	)

	if err == nil {
		t.Fatal(
			"expected closed file",
		)
	}
}
