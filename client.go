package oac

import "net/http"

type Client struct {
	HTTPClient *http.Client
	Accounts   Accounts
}

func NewClient(httpClient *http.Client) Client {
	return Client{
		HTTPClient: httpClient,
		Accounts: Accounts{
			HTTPClient: httpClient,
		},
	}
}
