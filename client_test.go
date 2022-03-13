package oac

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(http.DefaultClient)
	assert.NotNil(t, c)
}
