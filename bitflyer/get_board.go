package bitflyer

import (
	"context"
)

type Order struct {
	Price float64
	Size  float64
}

type GetBoardOutput struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

func (c *Client) GetBoard(ctx context.Context, productCode string) (*GetBoardOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "board", nil, nil, false)
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

	output := GetBoardOutput{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
