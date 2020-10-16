package bitflyer

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
)

const (
	BaseURL    = "https://api.bitflyer.com/"
	APIVersion = "v1"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Region      string
	BaseURL     *url.URL
	Credentials Credentials
	HTTPClient  HTTPClient
}

// NewClient returns a new bitFlyer API client.
func NewClient(cfg Config) *Client {
	u, _ := url.Parse(BaseURL)

	client := &Client{
		Region:  cfg.Region,
		BaseURL: u,
		Credentials: Credentials{
			APIKey:    cfg.Credentials.APIKey,
			APISecret: cfg.Credentials.APISecret,
		},
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client
}

// NewRequestPublic returns a http request for public API.
func (c *Client) NewRequestPublic(ctx context.Context, method, spath string, body []byte) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, APIVersion, spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRequestPrivate returns a http request for private API.
func (c *Client) NewRequestPrivate(ctx context.Context, method, spath string, body []byte) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, APIVersion, "me", spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	setAuthHeaders(req.Header, c.Credentials, method, u, body)

	return req, nil
}

func setAuthHeaders(header http.Header, credentials Credentials, method string, path url.URL, body []byte) {
	now := time.Now().Unix()
	timestamp := strconv.FormatInt(now, 10)

	h := hmac.New(sha256.New, []byte(credentials.APISecret))
	h.Write([]byte(timestamp))
	h.Write([]byte(method))
	h.Write([]byte(path.Path))
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
	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("BODY: %#v", string(b))
	decoder := json.NewDecoder(resp.Body)

	return decoder.Decode(out)
}
