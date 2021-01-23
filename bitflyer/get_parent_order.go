package bitflyer

import (
	"context"
	"fmt"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type Parameter struct {
	ProductCode   types.ProductCode `json:"product_code"`
	ConditionType string            `json:"condition_type"`
	Side          types.OrderSide   `json:"side"`
	Price         int64             `json:"price"`
	Size          float64           `json:"Size"`
	TriggerPrice  int64             `json:"trigger_price"`
	Offset        int64             `json:"offset"`
}

type ParentOrderDetails struct {
	ID                      int64        `json:"id"`
	ParentOrderID           string       `json:"parent_order_id"`
	OrderMethod             string       `json:"order_method"`
	ExpireDate              string       `json:"expire_date"`
	TimeInForce             string       `json:"time_in_force"`
	Parameters              []*Parameter `json:"parameters"`
	ParentOrderAcceptanceID string       `json:"parent_order_acceptance_id"`
}

// GetParentOrderOutput represents an output of GetParentOrderList method.
type GetParentOrderOutput struct {
	ParentOrderDetails ParentOrderDetails
}

// GetParentOrder gets parent orders.
func (c *Client) GetParentOrder(ctx context.Context, parentOrderID string, parentOrderAcceptanceID string) (*GetParentOrderOutput, error) {
	if len(parentOrderID) != 0 && len(parentOrderAcceptanceID) != 0 {
		return nil, fmt.Errorf(
			"do not input both parentOrderID and parentOrderAcceptanceID. parentOrpderID=%s, parentOrderAcceptanceID=%s",
			parentOrderID,
			parentOrderAcceptanceID,
		)
	}

	if len(parentOrderID) == 0 && len(parentOrderAcceptanceID) == 0 {
		return nil, fmt.Errorf(
			"input either parentOrderID or parentOrderAcceptanceID. parentOrderID=%s, parentOrderAcceptanceID=%s",
			parentOrderID,
			parentOrderAcceptanceID,
		)
	}

	req, err := c.NewRequest(ctx, "GET", "getparentorder", nil, nil, true)
	if err != nil {
		return nil, err
	}

	if len(parentOrderID) != 0 {
		q := req.URL.Query()
		q.Add("parent_order_id", parentOrderID)

		req.URL.RawQuery = q.Encode()
	}

	if len(parentOrderAcceptanceID) != 0 {
		q := req.URL.Query()
		q.Add("parent_order_acceptance_id", parentOrderAcceptanceID)

		req.URL.RawQuery = q.Encode()
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetParentOrderOutput{}
	if err := decodeBody(res, &output.ParentOrderDetails); err != nil {
		return nil, err
	}

	return &output, nil
}
