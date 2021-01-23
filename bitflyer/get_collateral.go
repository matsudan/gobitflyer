package bitflyer

import (
	"context"
)

type GetCollateralOutput struct {
	// The amount deposited in JPY.
	Collateral float64 `json:"collateral"`

	// The profit or loss from valuation.
	OpenPositionPnl float64 `json:"open_position_pnl"`

	// The current required margin.
	RequireCollateral float64 `json:"require_collateral"`

	// The current maintenance margin.
	KeepRate float64 `json:"keep_rate"`
}

func (c *Client) GetCollateral(ctx context.Context) (*GetCollateralOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getcollateral", nil, nil, true)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetCollateralOutput{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
