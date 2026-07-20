package taxii

import (
	"context"
	"testing"
)

func newTestClient(server *TestServer) *Client {

	return NewClient(Config{
		APIRoot:      server.Server.URL + "/taxii2/root",
		CollectionID: "collection-id",
		APIKey:       "test-token",
	})
}

func TestClientSend(t *testing.T) {

	server := NewTestServer(t)
	defer server.Close()

	client := newTestClient(server)

	err := client.Send(
		context.Background(),
		[]byte(`{"type":"bundle","objects":[]}`),
	)

	if err != nil {
		t.Fatal(err)
	}

	if !server.DiscoveryCalled {
		t.Fatal("Discovery endpoint not called")
	}

	if !server.CollectionsCalled {
		t.Fatal("Collections endpoint not called")
	}

	if !server.ObjectsCalled {
		t.Fatal("Objects endpoint not called")
	}
}

func TestClientMissingCollection(t *testing.T) {

	server := NewTestServer(t)
	server.MissingCollection = true
	defer server.Close()

	client := newTestClient(server)
	client.config.CollectionID = "missing"

	err := client.Send(
		context.Background(),
		[]byte(`{}`),
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestClientUnauthorized(t *testing.T) {

	server := NewTestServer(t)
	server.Unauthorized = true
	defer server.Close()

	client := newTestClient(server)

	err := client.Send(
		context.Background(),
		[]byte(`{}`),
	)

	if err == nil {
		t.Fatal("expected unauthorized error")
	}
}
