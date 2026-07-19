package taxii

import (
	"context"
)

type Sender struct {
	client *Client
}

func NewSender(
	config Config,
) *Sender {

	return &Sender{

		client: NewClient(config),
	}
}

func (s *Sender) Send(
	ctx context.Context,
	data []byte,
) error {

	return s.client.Send(
		ctx,
		data,
	)
}

func (s *Sender) Close() error {

	return nil
}
