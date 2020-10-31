package bitflyer

import (
	"context"
)

type Balance struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}

type GetBalanceOutput struct {
	Balance []*Balance
}

func (c *Client) GetBalanceList(ctx context.Context) (*GetBalanceOutput, error) {
	req, err := c.NewRequestPrivate(ctx, "GET", "getbalance", nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetBalanceOutput{}
	if err := decodeBody(res, &output.Balance); err != nil {
		return nil, err
	}

	return &output, nil
}
