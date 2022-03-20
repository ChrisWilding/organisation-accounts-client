package oac

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var name = uuid.New().String()
var attrs = AccountAttributes{
	BankID:       "400300",
	BankIDCode:   "GBDSC",
	BaseCurrency: "GBP",
	Bic:          "NWBKGB22",
	Country:      String("GB"),
	Name:         []string{name},
}

var acc = AccountData{
	Attributes:     &attrs,
	ID:             uuid.New().String(),
	OrganisationID: uuid.New().String(),
	Type:           "accounts",
}

func TestCreateAccount(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Create(context.Background(), &acc)

	assert.Equal(t, attrs.BankID, res.Attributes.BankID)
	assert.Equal(t, attrs.BankIDCode, res.Attributes.BankIDCode)
	assert.Equal(t, attrs.BaseCurrency, res.Attributes.BaseCurrency)
	assert.Equal(t, attrs.Country, res.Attributes.Country)
	assert.Len(t, attrs.Name, 1)
	assert.Equal(t, attrs.Name[0], res.Attributes.Name[0])
	assert.Equal(t, acc.ID, res.ID)
	assert.Equal(t, acc.OrganisationID, res.OrganisationID)
	assert.Equal(t, acc.Type, "accounts")
	assert.Equal(t, int64(0), *res.Version)

	assert.Nil(t, err)
}

func TestCreateAccountWithSameID(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Create(context.Background(), &acc)

	assert.Nil(t, res)
	assert.ErrorIs(t, err, ErrAccountIsDuplicate)
}

func TestCreateAccountWithMissingMandatoryFields(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	var name = uuid.New().String()
	var attrs = AccountAttributes{
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BaseCurrency: "GBP",
		Bic:          "NWBKGB22",
		Country:      nil,
		Name:         []string{name},
	}

	var acc = AccountData{
		Attributes:     &attrs,
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
	}

	res, err := c.Accounts.Create(context.Background(), &acc)

	assert.Nil(t, res)
	assert.ErrorIs(t, err, ErrAccountBadRequest)
}

func TestFetchAccount(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Fetch(context.Background(), acc.ID)

	assert.Equal(t, attrs.BankID, res.Attributes.BankID)
	assert.Equal(t, attrs.BankIDCode, res.Attributes.BankIDCode)
	assert.Equal(t, attrs.BaseCurrency, res.Attributes.BaseCurrency)
	assert.Equal(t, attrs.Country, res.Attributes.Country)
	assert.Len(t, attrs.Name, 1)
	assert.Equal(t, attrs.Name[0], res.Attributes.Name[0])
	assert.Equal(t, acc.ID, res.ID)
	assert.Equal(t, acc.OrganisationID, res.OrganisationID)
	assert.Equal(t, acc.Type, "accounts")
	assert.Equal(t, int64(0), *res.Version)

	assert.Nil(t, err)
}

func TestFetchAccountWhenNotFound(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Fetch(context.Background(), uuid.New().String())
	assert.Nil(t, res)
	assert.ErrorIs(t, err, ErrAccountNotFound)
}

func TestDeleteAccountWhenNotFound(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Delete(context.Background(), uuid.New().String(), 0)

	assert.False(t, res)
	assert.ErrorIs(t, err, ErrAccountNotFound)
}

func TestDeleteAccountWhenVersionIsIncorrect(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Delete(context.Background(), acc.ID, 999)

	assert.False(t, res)
	assert.ErrorIs(t, err, ErrAccountVersionIncorrect)
}

func TestDeleteAccount(t *testing.T) {
	c := NewClient(http.DefaultClient, WithBaseURL("http://localhost:8080"))

	res, err := c.Accounts.Delete(context.Background(), acc.ID, 0)

	assert.True(t, res)
	assert.Nil(t, err)
}
