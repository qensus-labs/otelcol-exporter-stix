package stixexporter

import (
	"io"
	"os"
)

func createWriter(
	output string,
) (io.Writer, io.Closer, error) {

	switch output {

	case "", "stdout":

		return os.Stdout, nil, nil

	default:

		file, err := os.OpenFile(
			output,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)

		if err != nil {
			return nil, nil, err
		}

		return file, file, nil
	}
}
