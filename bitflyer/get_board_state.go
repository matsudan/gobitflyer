package bitflyer

import (
	"context"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

type GetBoardStateOutput struct {
	Health types.Health `json:"health"`
	State  types.State  `json:"state"`
	Data   struct {
		SpecialQuotation float64 `json:"special_quotation"`
	} `json:"data"`
}

func (c *Client) GetBoardState(ctx context.Context, productCode string) (*GetBoardStateOutput, error) {
	req, err := c.NewRequestPublic(ctx, "GET", "getboardstate", nil)
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

	output := GetBoardStateOutput{}

	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
