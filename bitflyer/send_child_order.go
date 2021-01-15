package bitflyer

import (
	"context"
	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type SendChildOrderInput struct {
	ProductCode    types.ProductCode    `json:"product_code"`
	ChildOrderType types.ChildOrderType `json:"child_order_type"`
	Side           types.OrderSide      `json:"side"`
	Price          int64                `json:"price"`
	Size           float64              `json:"size"`
	MinuteToExpire int64                `json:"minute_to_expire"`
	TimeInForce    types.TimeInForce    `json:"time_in_force"`
}

type SendChildOrderOutput struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func (c *Client) SendChildOrder(ctx context.Context, sendChildOrderInput *SendChildOrderInput) (*SendChildOrderOutput, error) {

	req, err := c.NewRequestPrivate(ctx, "POST", "sendchildorder", sendChildOrderInput, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := SendChildOrderOutput{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
