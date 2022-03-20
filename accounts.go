package oac

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Accounts groups functionality to create, delete and fetch AccountData.
type Accounts struct {
	c *Client
}

type accountReq struct {
	Data *AccountData `json:"data,omitempty"`
}

type accountRes struct {
	Data *AccountData `json:"data,omitempty"`
}

// Create a new account using the provided AccountData.
func (a *Accounts) Create(ctx context.Context, accountData *AccountData) (*AccountData, error) {
	var ar = accountReq{
		Data: accountData,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(ar)
	if err != nil {
		return nil, err
	}

	url := a.c.baseURL + "/v1/organisation/accounts"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.api+json")

	res, err := a.c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var out accountRes
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out.Data, nil
}

// func (a *Accounts) Delete() (bool, error) {
// 	return true, nil
// }

// func (a *Accounts) Fetch() (*Account, error) {
// 	return nil, nil
// }
