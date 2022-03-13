package oac

import "net/http"

type Accounts struct {
	HTTPClient *http.Client //nolint
}

type Account struct {
}

func (a *Accounts) Create(acount Account) (*Account, error) {
	return nil, nil
}

func (a *Accounts) Delete() (bool, error) {
	return true, nil
}

func (a *Accounts) Fetch() (*Account, error) {
	return nil, nil
}
