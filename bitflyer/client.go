package bitflyer

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://api.bitflyer.com/"
	apiVersion     = "v1"
	userAgent      = "gobitflyer"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL     *url.URL
	Credentials Credentials
	HTTPClient  HTTPClient

	// User agent used when communicating with the bitFlyer Lightning API.
	UserAgent string
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
		UserAgent:  userAgent,
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

// NewRequest returns a http request.
func (c *Client) NewRequest(ctx context.Context, method, pathStr string, body interface{}, paginationQuery *PaginationQuery, isPrivate bool) (*http.Request, error) {
	u := *c.BaseURL
	if isPrivate {
		u.Path = path.Join(c.BaseURL.Path, apiVersion, "me", pathStr)
	} else {
		u.Path = path.Join(c.BaseURL.Path, apiVersion, pathStr)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if isPrivate {
		setAuthHeaders(req.Header, c.Credentials, method, u)
	}

	if paginationQuery == nil {
		return req, nil
	}

	paginationQuery.setPaginationQueries(req)

	return req, nil
}

func setAuthHeaders(header http.Header, credentials Credentials, method string, path url.URL) {
	now := time.Now().Unix()
	timestamp := strconv.FormatInt(now, 10)

	h := hmac.New(sha256.New, []byte(credentials.APISecret))
	h.Write([]byte(timestamp))
	h.Write([]byte(method))
	h.Write([]byte(path.Path))

	sign := hex.EncodeToString(h.Sum(nil))

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

func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		return resp, err
	}

	return resp, err
}

type Error struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	Data         string `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Status: %v, Message: %v, Data: %v",
		e.Status, e.ErrorMessage, e.Data)
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Err      Error
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d, error: %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Err.Error(),
	)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse.Err)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return errorResponse
}
