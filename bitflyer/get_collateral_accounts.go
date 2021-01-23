package bitflyer

import (
	"context"
)

type CollateralAccount struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

type GetCollateralAccountsOutput struct {
	CollateralAccounts []*CollateralAccount
}

func (c *Client) GetCollateralAccountList(ctx context.Context) (*GetCollateralAccountsOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getcollateralaccounts", nil, nil, true)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetCollateralAccountsOutput{}
	if err := decodeBody(res, &output.CollateralAccounts); err != nil {
		return nil, err
	}

	return &output, nil
}
