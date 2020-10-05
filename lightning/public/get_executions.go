package public

import "context"

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

type PaginationQuery struct {
	Count  string
	Before string
	After  string
}

func (c *Client) GetExecutionList(ctx context.Context, productCode string, paginationQuery PaginationQuery) (*GetExecutionListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "executions", nil)
	if err != nil {
		return nil, err
	}

	count := paginationQuery.Count
	before := paginationQuery.Before
	after := paginationQuery.After

	q := req.URL.Query()
	q.Add("product_code", productCode)

	if count != "" {
		q.Add("count", count)
	}
	if before != "" {
		q.Add("before", before)
	}
	if after != "" {
		q.Add("after", after)
	}

	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetExecutionListOutput{}
	if err := decodeBody(res, &output.Executions); err != nil {
		return nil, err
	}

	return &output, nil
}
