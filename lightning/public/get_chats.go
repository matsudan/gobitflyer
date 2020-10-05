package public

import "context"

type Chat struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

func (c *Client) GetChatList(ctx context.Context, fromDate string) ([]*Chat, error) {
	req, err := c.NewRequest(ctx, "GET", "getchats", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("from_date", fromDate)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	output := []*Chat{}
	if err := decodeBody(res, &output); err != nil {
		return nil, err
	}

	return output, nil
}
