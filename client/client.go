package client

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client HTTP client
type Client struct {
	base     string
	user     string
	password string
	req      *resty.Request
}

// New init client
func New(base, user, password string) Client {
	request := resty.New().SetDisableWarn(true).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		R().
		SetBasicAuth(user, password)

	client := Client{
		base:     base,
		user:     user,
		password: password,
		req:      request,
	}
	return client
}

// Login verify login
func (c Client) Login() error {
	resp, err := c.Get("/org", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New("Fail to login Grafana server with provided information")
	}
	return nil
}

// Get HTTP GET
func (c Client) Get(uri string, params map[string]string) (*resty.Response, error) {
	if params == nil {
		return c.req.Get(c.base + uri)
	}
	return c.req.SetQueryParams(params).Get(c.base + uri)
}

// Post HTTP POST
func (c Client) Post(uri string, payload map[string]interface{}) (*resty.Response, error) {
	return c.req.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}).SetBody(payload).Post(c.base + uri)
}

// Delete HTTP DELETE
func (c Client) Delete(uri string) (*resty.Response, error) {
	return c.req.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}).Delete(c.base + uri)
}

// timeToEpochMS transfer time into epoch datetime in milliseconds
func timeToEpochMS(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// epochMSToTime transfer epoch datetime in milliseconds to time (RFC3339)
func epochMSToTime(ms int64) string {
	t := time.Unix(0, ms*1000000)
	return t.Format(time.RFC3339)
}

// ParseRFC3339TimeString parse string to RFC3339 time
func ParseRFC3339TimeString(ts string) (time.Time, error) {
	return time.Parse(time.RFC3339, ts)
}
