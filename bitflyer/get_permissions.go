package bitflyer

import (
	"context"
)

type GetPermissionListOutput struct {
	Permissions []string
}

func (c *Client) GetPermissionList(ctx context.Context) (*GetPermissionListOutput, error) {
	req, err := c.NewRequestPrivate(ctx, "GET", "getpermissions", nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := GetPermissionListOutput{}
	if err := decodeBody(res, &output.Permissions); err != nil {
		return nil, err
	}

	return &output, nil
}
