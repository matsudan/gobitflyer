package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type ParentOrder struct {
	ID                      int64                  `json:"id"`
	ParentOrderID           string                 `json:"parent_order_id"`
	ProductCode             types.ProductCode      `json:"product_code"`
	Side                    types.OrderSide        `json:"side"`
	ParentOrderType         string                 `json:"parent_order_type"`
	Price                   int64                  `json:"price"`
	AveragePrice            int64                  `json:"average_price"`
	Size                    float64                `json:"size"`
	ParentOrderState        types.ParentOrderState `json:"parent_order_state"`
	ExpireDate              string                 `json:"expire_date"`
	ParentOrderDate         string                 `json:"parent_order_date"`
	ParentOrderAcceptanceID string                 `json:"parent_order_acceptance_id"`
	OutstandingSize         int64                  `json:"outstanding_size"`
	CancelSize              float64                `json:"cancel_size"`
	ExecutedSize            float64                `json:"executed_size"`
	TotalCommission         int64                  `json:"total_commission"`
}

// GetParentOrderListOutput represents an output of GetParentOrderList method.
type GetParentOrderListOutput struct {
	ParentOrders []*ParentOrder
}

// GetParentOrderList gets parent orders.
func (c *Client) GetParentOrderList(ctx context.Context, parentOrderState *types.ParentOrderState, paginationQuery *PaginationQuery) (*GetParentOrderListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getparentorders", nil, paginationQuery, true)
	if err != nil {
		return nil, err
	}

	if parentOrderState != nil {
		q := req.URL.Query()
		q.Add("parent_order_state", string(*parentOrderState))

		req.URL.RawQuery = q.Encode()
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetParentOrderListOutput{}
	if err := decodeBody(res, &output.ParentOrders); err != nil {
		return nil, err
	}

	return &output, nil
}
