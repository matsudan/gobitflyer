package lightning

import (
	"context"
)

type GetHealthOutput struct {
	Status string `json:"status"`
}

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

func (c *Client) GetHealth(ctx context.Context) (*GetHealthOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "gethealth", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var output GetHealthOutput
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (c *Client) GetMarketList(ctx context.Context) ([]*Market, error) {
	req, err := c.NewRequest(ctx, "GET", "getmarkets", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := []*Market{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}
