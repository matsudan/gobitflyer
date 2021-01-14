package bitflyer

import "github.com/matsudan/gobitflyer/bitflyer/types"

type ChildOrderRequest struct {
	ProductCode    types.ProductCode    `json:"product_code"`
	ChildOrderType types.ChildOrderType `json:"child_order_type"`
	Side           types.OrderSide      `json:"side"`
	Price          int64                `json:"price"`
	Size           float64              `json:"size"`
	MinuteToExpire int64                `json:"minute_to_expire"`
	TimeInForce    types.TimeInForce    `json:"time_in_force"`
}
