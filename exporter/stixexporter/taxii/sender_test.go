package taxii

import "testing"

func TestNewSender(t *testing.T) {

	sender := NewSender(
		Config{
			URL: "https://taxii.example.com/taxii2/",
		},
	)

	if sender == nil {
		t.Fatal(
			"expected sender",
		)
	}
}
