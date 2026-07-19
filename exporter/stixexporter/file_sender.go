package stixexporter

import (
	"context"
	"os"
)

type fileSender struct {
	file *os.File
}

func newFileSender(
	path string,
) (Sender, error) {

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &fileSender{
		file: file,
	}, nil
}

func (s *fileSender) Send(
	ctx context.Context,
	data []byte,
) error {

	_, err := s.file.Write(
		data,
	)

	return err
}

func (s *fileSender) Close() error {

	return s.file.Close()
}
