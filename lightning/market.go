package lightning

import (
	"context"
)

type GetHealthOutput struct {
	Status string `json:"status"`
}

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

type Order struct {
	Price float64
	Size  float64
}

type Board struct {
	MidPrice string  `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	TimeStamp       string  `json:"timestamp"`
	TickID          int64   `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	LTP             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (c *Client) GetHealth(ctx context.Context) (*GetHealthOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "gethealth", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var output GetHealthOutput
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (c *Client) GetMarketList(ctx context.Context) ([]*Market, error) {
	req, err := c.NewRequest(ctx, "GET", "markets", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := []*Market{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}

//func (c *Client) GetBoard(ctx context.Context, productCode string) (*Board, error) {
//	req, err := c.NewRequest(ctx, "GET", "board", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	q := req.URL.Query()
//	q.Add("product_code", productCode)
//	req.URL.RawQuery = q.Encode()
//
//	res, err := c.HTTPClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	output := Board{}
//	if err := decodeBody(res, &output); err != nil {
//		return nil, err
//	}
//
//	return &output, nil
//}

func (c *Client) GetTicker(ctx context.Context, productCode string) (*Ticker, error) {
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

	output := Ticker{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
