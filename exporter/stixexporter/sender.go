package stixexporter

import "context"

// Sender defines how STIX bundles are delivered.
type Sender interface {
	Send(
		ctx context.Context,
		data []byte,
	) error

	Close() error
}
