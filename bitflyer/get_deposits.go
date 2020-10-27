package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

// Deposit represents a deposit of your bitFlyer account.
type Deposit struct {
	ID           int64               `json:"id"`
	OrderID      string              `json:"order_id"`
	CurrencyCode string              `json:"currency_code"`
	Amount       int64               `json:"amount"`
	Status       types.DepositStatus `json:"status"`
	EventDate    string              `json:"event_date"`
}

// GetDepositListOutput represent an output of GetDepositList method.
type GetDepositListOutput struct {
	Deposits []*Deposit
}

// GetDepositList gets cash deposits.
func (c *Client) GetDepositList(ctx context.Context) (*GetDepositListOutput, error) {
	req, err := c.NewRequestPrivate(ctx, "GET", "getdeposits", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetDepositListOutput{}
	if err := decodeBody(res, &output.Deposits); err != nil {
		return nil, err
	}

	return &output, nil
}
