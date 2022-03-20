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

func TestOverrideClientBaseURL(t *testing.T) {
	expected := "https://www.example.com"
	c := NewClient(http.DefaultClient, WithBaseURL(expected))
	assert.Equal(t, expected, c.baseURL)
}
