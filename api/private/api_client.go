package private

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/matsudan/gobitflyer/lightning"
)

const (
	BaseURL = "https://api.bitflyer.com/v1"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Region      string
	EndpointURL *url.URL
	Credentials lightning.Credentials
	HTTPClient  HTTPClient
}

func New() (*Client, error) {
	u, err := url.Parse(BaseURL)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		EndpointURL: u,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client, nil
}

func NewFromConfig(cfg lightning.Config) (*Client, error) {
	u, err := url.Parse(BaseURL)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		Region:      cfg.Region,
		EndpointURL: u,
		Credentials: lightning.Credentials{
			APIKey:    cfg.Credentials.APIKey,
			APISecret: cfg.Credentials.APISecret,
		},
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client, nil
}

func (c *Client) NewRequest(ctx context.Context, method, spath string, body []byte) (*http.Request, error) {
	u := *c.EndpointURL
	u.Path = path.Join(c.EndpointURL.Path, spath)

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	setAuthHeaders(req.Header, c.Credentials, method, u.String(), body)

	req = req.WithContext(ctx)

	return req, nil
}

func setAuthHeaders(header http.Header, credentials lightning.Credentials, method string, path string, body []byte) {
	now := time.Now().Unix()
	timestamp := strconv.FormatInt(now, 10)

	h := hmac.New(sha256.New, []byte(credentials.APISecret))
	h.Write([]byte(timestamp))
	h.Write([]byte(method))
	h.Write([]byte(path))
	if len(body) > 0 {
		h.Write(body)
	}
	sign := hex.EncodeToString(h.Sum(nil))

	header.Set("Content-Type", "application/json")
	header.Set("ACCESS-KEY", credentials.APIKey)
	header.Set("ACCESS-TIMESTAMP", timestamp)
	header.Set("ACCESS-SIGN", sign)
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	return decoder.Decode(out)
}
