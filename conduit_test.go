package gonduit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/test/server"
)

func TestConduitQuery(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	response := server.ResponseFromJSON(`{
	  "result": {
		"phid.query": {
		  "description": "Retrieve information about arbitrary PHIDs.",
		  "params": {
			"phids": "required list<phid>"
		  },
		  "return": "nonempty dict<string, wild>"
		}
	  }
	}`)
	s.RegisterMethod("conduit.query", 200, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)

	r, err := c.ConduitQuery()
	assert.Nil(t, err)

	assert.Equal(
		t,
		"nonempty dict<string, wild>",
		(*r)["phid.query"].Return,
	)
}
