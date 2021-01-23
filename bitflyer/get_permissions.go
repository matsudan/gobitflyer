package bitflyer

import (
	"context"
)

type GetPermissionListOutput struct {
	Permissions []string
}

func (c *Client) GetPermissionList(ctx context.Context) (*GetPermissionListOutput, error) {
	req, err := c.NewRequest(ctx, "GET", "getpermissions", nil, nil, true)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetPermissionListOutput{}
	if err := decodeBody(res, &output.Permissions); err != nil {
		return nil, err
	}

	return &output, nil
}
