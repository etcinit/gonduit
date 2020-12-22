package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeHTTPClient(t *testing.T) {
	client := makeHTTPClient(&ClientOptions{
		InsecureSkipVerify: true,
	})

	client2 := makeHTTPClient(&ClientOptions{
		InsecureSkipVerify: false,
	})

	assert.NotNil(t, client)
	assert.NotNil(t, client2)

	// Test that a provided client is used instead of a new one
	// with default settings.
	assert.Equal(t, client, makeHTTPClient(&ClientOptions{
		Client: client,
	}))
}
