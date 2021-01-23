package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type ChildOrder struct {
	ID                     int64                 `json:"id"`
	ChildOrderID           string                `json:"child_order_id"`
	ProductCode            types.ProductCode     `json:"product_code"`
	Side                   types.OrderSide       `json:"side"`
	ChildOrderType         types.ChildOrderType  `json:"child_order_type"`
	Price                  int64                 `json:"price"`
	AveragePrice           int64                 `json:"average_price"`
	Size                   float64               `json:"size"`
	ChildOrderState        types.ChildOrderState `json:"child_order_state"`
	ExpireDate             string                `json:"expire_date"`
	ChildOrderDate         string                `json:"child_order_date"`
	ChildOrderAcceptanceID string                `json:"child_order_acceptance_id"`
	OutstandingSize        int64                 `json:"outstanding_size"`
	CancelSize             float64               `json:"cancel_size"`
	ExecutedSize           float64               `json:"executed_size"`
	TotalCommission        int64                 `json:"total_commission"`
}

// GetChildOrderListOutput represent an output of GetChildOrderList method.
type GetChildOrderListOutput struct {
	ChildOrders []*ChildOrder
}

// GetChildOrderList gets child orders.
func (c *Client) GetChildOrderList(ctx context.Context, paginationQuery *PaginationQuery) (*GetChildOrderListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getchildorders", nil, paginationQuery, true)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetChildOrderListOutput{}
	if err := decodeBody(res, &output.ChildOrders); err != nil {
		return nil, err
	}

	return &output, nil
}
