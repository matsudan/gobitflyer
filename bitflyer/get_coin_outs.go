package bitflyer

import (
	"context"
)

type CoinOut struct {
	ID            string  `json:"id"`
	OrderID       string  `json:"order_id"`
	CurrencyCode  string  `json:"currency_code"`
	Amount        float64 `json:"amount"`
	Address       string  `json:"address"`
	TxHash        string  `json:"tx_hash"`
	Fee           float64 `json:"fee"`
	AdditionalFee float64 `json:"additional_fee"`
	Status        string  `json:"status"`
	EventDate     string  `json:"event_date"`
}

type GetCoinOutListOutput struct {
	CoinOuts []*CoinOut
}

func (c *Client) GetCoinOutList(ctx context.Context) (*GetCoinOutListOutput, error) {
	// TODO: add pagination
	req, err := c.NewRequestPrivate(ctx, "GET", "getcoinouts", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetCoinOutListOutput{}
	if err := decodeBody(res, &output.CoinOuts); err != nil {
		return nil, err
	}

	return &output, nil
}
