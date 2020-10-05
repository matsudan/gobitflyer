package public

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	BaseURL = "https://api.bitflyer.com"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	EndpointURL *url.URL
	HTTPClient  HTTPClient
}

type Options struct {
	// Lightning api version.
	Version string
}

func New(options Options) (*Client, error) {
	u, err := url.Parse(BaseURL)
	if err != nil {
		return &Client{}, err
	}
	u.Path = path.Join(u.Path, options.Version)

	client := &Client{
		EndpointURL: u,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client, nil
}

func (c *Client) NewRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.EndpointURL
	u.Path = path.Join(c.EndpointURL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	return decoder.Decode(out)
}
