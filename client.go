package oac

import (
	"net/http"
)

const BASE_URL = "https://api.form3.tech"

// Client is the root of the API.
type Client struct {
	HTTPClient *http.Client
	Accounts   *Accounts
	baseURL    string
}

// ClientOption is used to customise the Client.
type ClientOption func(c *Client)

// WithBaseURL returns a ClientOption that changes the API base URL used by the Client.
// For example, this can be used to override the base URL for tests or to use a non-production
// environment.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient initialised a Client, wrapping an implementation of http.Client and a slice
// of ClientOptions that are used to customise the Client.
func NewClient(httpClient *http.Client, opts ...ClientOption) Client {
	c := Client{
		HTTPClient: httpClient,
	}
	c.Accounts = &Accounts{
		c: &c,
	}
	c.baseURL = BASE_URL

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

// String is a utility function to convert a string value to a pointer.
func String(s string) *string {
	return &s
}
