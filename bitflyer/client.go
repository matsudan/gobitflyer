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
	"os"
	"path"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://api.bitflyer.com/"
	APIVersion     = "v1"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL     *url.URL
	Credentials Credentials
	HTTPClient  HTTPClient
}

// NewClient returns a new bitFlyer API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Minute,
		}
	}

	u, _ := url.Parse(defaultBaseURL)

	apiKey := os.Getenv("BITFLYER_API_KEY")
	apiSecret := os.Getenv("BITFLYER_API_SECRET")

	client := &Client{
		BaseURL: u,
		Credentials: Credentials{
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
		HTTPClient: httpClient,
	}

	return client
}

type PaginationQuery struct {
	Count  string
	Before string
	After  string
}

func (p *PaginationQuery) setPaginationQueries(req *http.Request) {
	q := req.URL.Query()

	if p.Count != "" {
		q.Add("count", p.Count)
	}
	if p.Before != "" {
		q.Add("before", p.Before)
	}
	if p.After != "" {
		q.Add("after", p.After)
	}

	req.URL.RawQuery = q.Encode()
}

// NewRequestPublic returns a http request for public API.
func (c *Client) NewRequestPublic(ctx context.Context, method, spath string, body []byte, paginationQuery *PaginationQuery) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, APIVersion, spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if paginationQuery == nil {
		return req, nil
	}

	paginationQuery.setPaginationQueries(req)

	return req, nil
}

// NewRequestPrivate returns a http request for private API.
func (c *Client) NewRequestPrivate(ctx context.Context, method, spath string, body []byte, paginationQuery *PaginationQuery) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, APIVersion, "me", spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	setAuthHeaders(req.Header, c.Credentials, method, u, body)

	if paginationQuery == nil {
		return req, nil
	}

	paginationQuery.setPaginationQueries(req)

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
