package bitflyer

import (
	"context"
)

type CoinIn struct {
	ID           string  `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Address      string  `json:"address"`
	TxHash       string  `json:"tx_hash"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

type GetCoinInListOutput struct {
	CoinIns []*CoinIn
}

func (c *Client) GetCoinInList(ctx context.Context, paginationQuery *PaginationQuery) (*GetCoinInListOutput, error) {
	// TODO: add pagination
	req, err := c.NewRequest(ctx, "GET", "getcoinins", nil, paginationQuery, true)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetCoinInListOutput{}
	if err := decodeBody(res, &output.CoinIns); err != nil {
		return nil, err
	}

	return &output, nil
}
