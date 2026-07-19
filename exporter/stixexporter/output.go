package stixexporter

import (
	"io"
	"os"
)

func createWriter(output string) (io.Writer, error) {

	switch output {

	case "", "stdout":
		return os.Stdout, nil

	default:
		return nil, nil
	}
}
