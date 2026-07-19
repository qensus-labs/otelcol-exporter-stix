package taxii

import "testing"

func TestNewSender(t *testing.T) {

	sender := NewSender(
		Config{
			APIRoot: "https://taxii.example.com/taxii2/root",

			CollectionID: "12345678-abcd-1234-abcd-123456789",

			APIKey: "test-token",
		},
	)

	if sender == nil {
		t.Fatal(
			"expected sender",
		)
	}
}

func TestConfigEndpoints(t *testing.T) {

	config := Config{
		APIRoot: "https://taxii.example.com/taxii2/root",

		CollectionID: "collection-id",
	}

	if got := config.DiscoveryEndpoint(); got !=
		"https://taxii.example.com/taxii2/root/discovery" {

		t.Fatalf(
			"unexpected discovery endpoint: %s",
			got,
		)
	}

	if got := config.CollectionsEndpoint(); got !=
		"https://taxii.example.com/taxii2/root/collections" {

		t.Fatalf(
			"unexpected collections endpoint: %s",
			got,
		)
	}

	if got := config.ObjectsEndpoint(); got !=
		"https://taxii.example.com/taxii2/root/collections/collection-id/objects/" {

		t.Fatalf(
			"unexpected objects endpoint: %s",
			got,
		)
	}
}
