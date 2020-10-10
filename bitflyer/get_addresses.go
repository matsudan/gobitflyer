package bitflyer

import (
	"context"
)

type Address struct {
	Type         string `json:"type"`
	CurrencyCode string `json:"currency_code"`
	Address      string `json:"address"`
}

type GetAddressListOutput struct {
	Addresses []*Address
}

func (c *Client) GetAddressList(ctx context.Context) (*GetAddressListOutput, error) {
	req, err := c.NewRequestPrivate(ctx, "GET", "getaddresses", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetAddressListOutput{}
	if err := decodeBody(res, &output.Addresses); err != nil {
		return nil, err
	}

	return &output, nil
}