package bitflyer

import (
	"context"
)

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

type GetMarketListOutput struct {
	Markets []*Market
}

func (c *Client) GetMarketList(ctx context.Context) (*GetMarketListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "markets", nil, nil, true)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetMarketListOutput{}
	if err := decodeBody(res, &output.Markets); err != nil {
		return nil, err
	}

	return &output, nil
}
