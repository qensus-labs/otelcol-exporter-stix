package taxii

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) Collections(
	ctx context.Context,
) (*Collections, error) {

	resp, err := c.doRequest(
		ctx,
		"GET",
		c.config.CollectionsEndpoint(),
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"unexpected status %d",
			resp.StatusCode,
		)
	}

	var collections Collections

	if err := json.NewDecoder(
		resp.Body,
	).Decode(&collections); err != nil {

		return nil, err
	}

	return &collections, nil
}
