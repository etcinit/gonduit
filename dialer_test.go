package gonduit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

func TestDial(t *testing.T) {
	s := server.New()
	defer s.Close()

	s.RegisterCapabilities()

	_, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})

	assert.Nil(t, err)
}

func TestDial_withInvalid(t *testing.T) {
	s := server.New()
	defer s.Close()

	s.RegisterMethod("conduit.getcapabilities", 200, server.ResponseFromJSON(`{
		"fake": "fake"
	}`))

	_, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})

	assert.NotNil(t, err)
}

func TestAssertSupportedCapabilities(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"urlencoded"},
		Output: []string{"json"},
	}

	assert.Nil(t, assertSupportedCapabilities(response, &core.ClientOptions{}))
}

func TestAssertSupportedCapabilities_withMissingInput(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"fake"},
		Output: []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrURLEncodedInputUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{}),
	)
}

func TestAssertSupportedCapabilities_withMissingOutput(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Input:  []string{"urlencoded"},
		Output: []string{"fake"},
	}

	assert.Equal(
		t,
		core.ErrJSONOutputUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{}),
	)
}

func TestAssertSupportedCapabilities_withNoToken(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Authentication: []string{"session"},
		Input:          []string{"urlencoded"},
		Output:         []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrTokenAuthUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{
			APIToken: "super-secret-token",
		}),
	)
}

func TestAssertSupportedCapabilities_withNoCertificate(t *testing.T) {
	response := responses.ConduitCapabilitiesResponse{
		Authentication: []string{"token"},
		Input:          []string{"urlencoded"},
		Output:         []string{"json"},
	}

	assert.Equal(
		t,
		core.ErrSessionAuthUnsupported,
		assertSupportedCapabilities(response, &core.ClientOptions{
			Cert:     "super-secret-token",
			CertUser: "alice",
		}),
	)
}
