package oac

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var acceptHeader = "application/vnd.api+json"

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
	req.Header.Set("Accept", acceptHeader)

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

var ErrAccountNotFound = errors.New("account does not exist")
var ErrAccountVersionIncorrect = errors.New("account version incorrect")

func (a *Accounts) Delete(ctx context.Context, id string, version int) (bool, error) {
	url := a.c.baseURL + "/v1/organisation/accounts/" + id
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	q := req.URL.Query()
	q.Add("version", strconv.Itoa(version))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", acceptHeader)

	res, err := a.c.HTTPClient.Do(req)
	if err != nil {
		return false, err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return false, ErrAccountNotFound
	case http.StatusConflict:
		return false, ErrAccountVersionIncorrect
	default:
		return res.StatusCode == http.StatusNoContent, nil
	}
}

func (a *Accounts) Fetch(ctx context.Context, id string) (*AccountData, error) {
	url := a.c.baseURL + "/v1/organisation/accounts/" + id
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", acceptHeader)

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
