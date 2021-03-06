package bitflyer

import (
	"context"
)

type Execution struct {
	ID                         int64   `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

type GetExecutionListOutput struct {
	Executions []*Execution
}

func (c *Client) GetExecutionList(ctx context.Context, productCode string, paginationQuery *PaginationQuery) (*GetExecutionListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "executions", nil, paginationQuery, false)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("product_code", productCode)

	req.URL.RawQuery = q.Encode()

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetExecutionListOutput{}
	if err := decodeBody(res, &output.Executions); err != nil {
		return nil, err
	}

	return &output, nil
}
