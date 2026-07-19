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

		file, err := os.OpenFile(
			output,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)

		if err != nil {
			return nil, err
		}

		return file, nil
	}
}
