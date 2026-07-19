package stixexporter

import (
	"fmt"

	"github.com/qensus-labs/otelcol-exporter-stix/exporter/stixexporter/taxii"
)

func createSender(
	cfg *Config,
) (Sender, error) {

	switch cfg.Mode {

	case "", "stdout":

		return newStdoutSender(), nil

	case "file":

		sender, err := newFileSender(
			cfg.Output,
		)

		if err != nil {
			return nil, err
		}

		return sender, nil

	case "taxii":

		return taxii.NewSender(
			cfg.TAXII,
		), nil

	default:

		return nil, fmt.Errorf(
			"unsupported STIX sender mode: %s",
			cfg.Mode,
		)
	}
}
