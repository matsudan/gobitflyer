package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type GetHealthOutput struct {
	Status types.ExchangeStatus `json:"status"`
}

func (c *Client) GetHealth(ctx context.Context) (*GetHealthOutput, error) {
	req, err := c.NewRequestPublic(ctx, "GET", "gethealth", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetHealthOutput{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}