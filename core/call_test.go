package core

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

func TestGetEndpointURI(t *testing.T) {
	assert.Equal(
		t,
		"phabricator.gonduit.wow/api/conduit.connect",
		GetEndpointURI("phabricator.gonduit.wow/", "conduit.connect"),
	)
}

func TestPerformCall(t *testing.T) {
	ts := server.New()
	defer ts.Close()
	ts.RegisterCapabilities()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.GetURL()+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Nil(t, err)
}

func TestPerformCall_withEmptyArray(t *testing.T) {
	ts := server.New()
	defer ts.Close()
	response := server.ResponseFromJSON(`{"result":[]}`)
	ts.RegisterMethod("phid.lookup", 200, response)

	var result responses.PHIDLookupResponse
	ptr := &result

	err := PerformCall(
		ts.GetURL()+"/api/phid.lookup",
		map[string]interface{}{},
		&ptr,
		&ClientOptions{},
	)

	assert.Nil(t, err)
}

func TestPerformCall_withErrorCode(t *testing.T) {
	ts := server.New()
	defer ts.Close()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.GetURL()+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	code := strconv.Itoa(http.StatusNotFound)
	assert.Equal(t, &ConduitError{
		code: code,
		info: "404 page not found",
	}, err)
}

func TestPerformCall_withBadHTTPResponseCode(t *testing.T) {
	ts := server.New()
	defer ts.Close()

	response := `{
	  "result": "Some result",
	  "error_code": "ERR-CONDUIT-CORE",
	  "error_info": "Something bad happened"
	}`
	ts.RegisterMethod("return.error", http.StatusOK, server.ResponseFromJSON(response))

	result := map[string]interface{}{}

	err := PerformCall(
		ts.GetURL()+"/api/return.error",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Equal(t, &ConduitError{
		code: "ERR-CONDUIT-CORE",
		info: "Something bad happened",
	}, err)
}

func TestPerformCall_withMissingResults(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	result := map[string]interface{}{}

	err := PerformCall(
		ts.URL+"/api/conduit.getcapabilities",
		map[string]interface{}{},
		&result,
		&ClientOptions{},
	)

	assert.Equal(t, ErrMissingResults, err)
}
