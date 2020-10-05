package public

import (
	"context"
)

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

func (c *Client) GetMarketList(ctx context.Context) ([]*Market, error) {
	req, err := c.NewRequest(ctx, "GET", "markets", nil)
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
