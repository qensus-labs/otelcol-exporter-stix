package stixexporter

import (
	"context"
	"os"
)

type stdoutSender struct {
}

func newStdoutSender() Sender {

	return &stdoutSender{}
}

func (s *stdoutSender) Send(
	ctx context.Context,
	data []byte,
) error {

	_, err := os.Stdout.Write(
		data,
	)

	return err
}

func (s *stdoutSender) Close() error {

	return nil
}
