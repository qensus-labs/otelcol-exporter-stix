package stixexporter

import (
	"testing"

	"github.com/qensus-labs/otelcol-exporter-stix/exporter/stixexporter/taxii"
)

func TestCreateSenderStdout(t *testing.T) {

	cfg := &Config{
		Mode: "stdout",
	}

	sender, err := createSender(cfg)

	if err != nil {
		t.Fatal(err)
	}

	if sender == nil {
		t.Fatal(
			"expected sender",
		)
	}
}

func TestCreateSenderTaxii(t *testing.T) {

	cfg := &Config{

		Mode: "taxii",

		TAXII: taxii.Config{
			APIRoot: "https://taxii.example.com/taxii2/root",

			CollectionID: "collection-id",
		},
	}

	sender, err := createSender(cfg)

	if err != nil {
		t.Fatal(err)
	}

	if sender == nil {
		t.Fatal(
			"expected TAXII sender",
		)
	}
}
