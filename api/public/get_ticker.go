package public

import (
	"context"
	"github.com/matsudan/gobitflyer/api/public/types"
)

type GetTickerOutput struct {
	ProductCode   string      `json:"product_code"`
	State         types.State `json:"state"`
	Timestamp     string      `json:"timestamp"`
	TickID        int64       `json:"tick_id"`
	BestBid       float64     `json:"best_bid"`
	BestAsk       float64     `json:"best_ask"`
	BestBidSize   float64     `json:"best_bid_size"`
	BestAskSize   float64     `json:"best_ask_size"`
	TotalBidDepth float64     `json:"total_bid_depth"`
	TotalAskDepth float64     `json:"total_ask_depth"`
	MarketBidSize float64     `json:"market_bid_size"`
	MarketAskSize float64     `json:"market_ask_size"`
	LTP             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (c *Client) GetTicker(ctx context.Context, productCode string) (*GetTickerOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "ticker", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("product_code", productCode)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetTickerOutput{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
