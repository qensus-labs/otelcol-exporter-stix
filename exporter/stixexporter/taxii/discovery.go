package taxii

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) Discovery(
	ctx context.Context,
) (*Discovery, error) {

	resp, err := c.doRequest(
		ctx,
		"GET",
		c.config.DiscoveryEndpoint(),
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

	var discovery Discovery

	if err := json.NewDecoder(
		resp.Body,
	).Decode(&discovery); err != nil {

		return nil, err
	}

	return &discovery, nil
}
