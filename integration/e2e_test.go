//go:build integration

// Package integration runs an end-to-end test against the stack in
// docker-compose.yaml: OTLP logs -> otelcol-stix -> Medallion (TAXII 2.1).
package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"
)

const (
	otlpLogsURL = "http://localhost:4318/v1/logs"

	taxiiObjectsURL = "http://localhost:9088/trustgroup1/collections/91a7b528-80eb-42ed-a74d-c6fbd5a26116/objects/"

	otlpPayload = `{
  "resourceLogs": [{
    "scopeLogs": [{
      "logRecords": [{
        "timeUnixNano": "%d",
        "body": {"stringValue": "e2e test"},
        "attributes": [
          {"key": "client.address", "value": {"stringValue": "203.0.113.10"}},
          {"key": "server.address", "value": {"stringValue": "example.org"}},
          {"key": "network.peer.address", "value": {"stringValue": "2001:db8::1"}},
          {"key": "url.full", "value": {"stringValue": "%s"}}
        ]
      }]
    }]
  }]
}`
)

type stixObject struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"`
	Value      string   `json:"value"`
	ObjectRefs []string `json:"object_refs"`
}

func TestLogsReachTAXIIServer(t *testing.T) {

	// Unique per run, so objects left over from earlier runs can't satisfy the test.
	url := fmt.Sprintf("https://example.org/e2e/%d", time.Now().UnixNano())

	sendLogs(t, fmt.Sprintf(otlpPayload, time.Now().UnixNano(), url))

	var urlObject stixObject

	eventually(t, func() bool {
		var ok bool
		urlObject, ok = findObject(fetchObjects(t), "url", url)
		return ok
	})

	objects := fetchObjects(t)

	for _, want := range []stixObject{
		{Type: "ipv4-addr", Value: "203.0.113.10"},
		{Type: "domain-name", Value: "example.org"},
		{Type: "ipv6-addr", Value: "2001:db8::1"},
	} {
		if _, ok := findObject(objects, want.Type, want.Value); !ok {
			t.Errorf("no %s object with value %q in collection", want.Type, want.Value)
		}
	}

	observed := slices.ContainsFunc(objects, func(o stixObject) bool {
		return o.Type == "observed-data" && slices.Contains(o.ObjectRefs, urlObject.ID)
	})

	if !observed {
		t.Errorf("no observed-data object references %s", urlObject.ID)
	}
}

// sendLogs posts OTLP/JSON logs, retrying while the Collector starts up.
func sendLogs(t *testing.T, payload string) {

	t.Helper()

	eventually(t, func() bool {
		resp, err := http.Post(otlpLogsURL, "application/json", strings.NewReader(payload))
		if err != nil {
			return false
		}
		defer resp.Body.Close()
		return resp.StatusCode == http.StatusOK
	})
}

func fetchObjects(t *testing.T) []stixObject {

	t.Helper()

	req, err := http.NewRequest(http.MethodGet, taxiiObjectsURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Accept", "application/taxii+json;version=2.1")
	req.SetBasicAuth("admin", "Password01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("GET objects: %s: %s", resp.Status, body)
	}

	var envelope struct {
		Objects []stixObject `json:"objects"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		t.Fatal(err)
	}

	return envelope.Objects
}

func findObject(objects []stixObject, objectType, value string) (stixObject, bool) {

	for _, o := range objects {
		if o.Type == objectType && o.Value == value {
			return o, true
		}
	}

	return stixObject{}, false
}

func eventually(t *testing.T, condition func() bool) {

	t.Helper()

	deadline := time.Now().Add(30 * time.Second)

	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("condition not met within 30s")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
