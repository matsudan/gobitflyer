package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

// Withdrawal represents a withdrawal from your bitFlyer account.
type Withdrawal struct {
	ID           int64                  `json:"id"`
	OrderID      string                 `json:"order_id"`
	CurrencyCode string                 `json:"currency_code"`
	Amount       int64                  `json:"amount"`
	Status       types.WithdrawalStatus `json:"status"`
	EventDate    string                 `json:"event_date"`
}

// GetWithdrawalListOutput represent an output of GetWithdrawalList method.
type GetWithdrawalListOutput struct {
	Withdrawals []*Withdrawal
}

// GetWithdrawalList gets withdrawal history.
func (c *Client) GetWithdrawalList(ctx context.Context, paginationQuery *PaginationQuery) (*GetWithdrawalListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getwithdrawals", nil, paginationQuery, true)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetWithdrawalListOutput{}
	if err := decodeBody(res, &output.Withdrawals); err != nil {
		return nil, err
	}

	return &output, nil
}
