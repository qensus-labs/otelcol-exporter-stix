package taxii

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	ContentTypeTAXII21 = "application/taxii+json;version=2.1"
)

type Client struct {
	config Config

	httpClient *http.Client
}

func NewClient(
	config Config,
) *Client {

	return &Client{
		config: config,

		httpClient: &http.Client{},
	}
}

// doRequest creates and executes a TAXII HTTP request.
func (c *Client) doRequest(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
) (*http.Response, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		ContentTypeTAXII21,
	)

	if c.config.APIKey != "" {

		req.Header.Set(
			"Authorization",
			"Bearer "+c.config.APIKey,
		)

	} else if c.config.Username != "" {

		req.SetBasicAuth(
			c.config.Username,
			c.config.Password,
		)
	}

	return c.httpClient.Do(req)
}

// validate verifies that the configured TAXII server
// and collection are reachable.
func (c *Client) validate(
	ctx context.Context,
) error {

	_, err := c.Discovery(ctx)

	if err != nil {
		return err
	}

	collections, err := c.Collections(ctx)

	if err != nil {
		return err
	}

	for _, collection := range collections.Collections {

		if collection.ID ==
			c.config.CollectionID {

			return nil
		}
	}

	return fmt.Errorf(
		"TAXII collection %q not found",
		c.config.CollectionID,
	)
}

// Send uploads a STIX bundle to the configured
// TAXII Objects endpoint.
func (c *Client) Send(
	ctx context.Context,
	data []byte,
) error {

	if err := c.validate(ctx); err != nil {
		return err
	}

	resp, err := c.doRequest(
		ctx,
		http.MethodPost,
		c.config.ObjectsEndpoint(),
		bytes.NewReader(data),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {

		return fmt.Errorf(
			"TAXII request failed: %s",
			resp.Status,
		)
	}

	return nil
}
