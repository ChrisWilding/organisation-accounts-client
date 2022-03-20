package oac

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	name := uuid.New().String()
	attrs := AccountAttributes{
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BaseCurrency: "GBP",
		Bic:          "NWBKGB22",
		Country:      String("GB"),
		Name:         []string{name},
	}

	acc := AccountData{
		Attributes:     &attrs,
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
	}

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
