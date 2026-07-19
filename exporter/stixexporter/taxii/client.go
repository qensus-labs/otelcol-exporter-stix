package taxii

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
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

func (c *Client) Send(
	ctx context.Context,
	data []byte,
) error {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.config.URL,
		bytes.NewReader(data),
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"Content-Type",
		"application/taxii+json;version=2.1",
	)

	if c.config.Username != "" {

		req.SetBasicAuth(
			c.config.Username,
			c.config.Password,
		)
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 ||
		resp.StatusCode >= 300 {

		return fmt.Errorf(
			"TAXII request failed: %s",
			resp.Status,
		)
	}

	return nil
}
