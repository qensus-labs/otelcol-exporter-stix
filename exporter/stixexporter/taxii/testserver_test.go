package taxii

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestServer struct {
	Server *httptest.Server

	DiscoveryCalled   bool
	CollectionsCalled bool
	ObjectsCalled     bool

	MissingCollection bool
	Unauthorized      bool
}

func NewTestServer(t *testing.T) *TestServer {

	ts := &TestServer{}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/taxii2/root/discovery",
		func(w http.ResponseWriter, r *http.Request) {

			ts.DiscoveryCalled = true

			if ts.Unauthorized {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Header().Set(
				"Content-Type",
				ContentTypeTAXII21,
			)

			_ = json.NewEncoder(w).Encode(
				Discovery{
					Title: "Test TAXII Server",
					APIRoots: []string{
						"https://example.com/taxii2/root",
					},
				},
			)
		},
	)

	mux.HandleFunc(
		"/taxii2/root/collections",
		func(w http.ResponseWriter, r *http.Request) {

			ts.CollectionsCalled = true

			w.Header().Set(
				"Content-Type",
				ContentTypeTAXII21,
			)

			collections := Collections{}

			if !ts.MissingCollection {
				collections.Collections = []Collection{
					{
						ID:    "collection-id",
						Title: "Default",
					},
				}
			}

			_ = json.NewEncoder(w).Encode(collections)
		},
	)

	mux.HandleFunc(
		"/taxii2/root/collections/collection-id/objects/",
		func(w http.ResponseWriter, r *http.Request) {

			ts.ObjectsCalled = true

			if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Fatalf("unexpected Authorization header: %q", got)
			}

			body, err := io.ReadAll(r.Body)

			if err != nil {
				t.Fatal(err)
			}

			if len(body) == 0 {
				t.Fatal("expected POST body")
			}

			w.WriteHeader(http.StatusAccepted)
		},
	)

	ts.Server = httptest.NewServer(mux)

	return ts
}

func (ts *TestServer) Close() {
	ts.Server.Close()
}
