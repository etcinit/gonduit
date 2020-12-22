package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/requests"
)

func TestMakeRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}

	options := &ClientOptions{
		APIToken: "hello-world-hello-world",
	}

	body, _ := prepareForm(&request, options)

	req, err := MakeRequest(
		"http://localhost/api/conduit.connect",
		&request,
		options,
	)

	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/api/conduit.connect", req.URL.String())
	assert.Equal(t, "POST", req.Method)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(req.Body)
	assert.Equal(t, body.Encode(), buffer.String())
}

func TestMakeRequest_withInvalidBody(t *testing.T) {
	options := &ClientOptions{
		APIToken: "hello-world-hello-world",
	}

	_, err := MakeRequest(
		"http://localhost/api/conduit.connect",
		http.Request{},
		options,
	)

	assert.NotNil(t, err)
}

func TestMakeRequest_withInvalidRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}

	options := &ClientOptions{
		APIToken: "hello-world-hello-world",
	}

	_, err := MakeRequest("htcp.\\invaid#&^%", &request, options)

	assert.NotNil(t, err)
}

func TestSetHeaders(t *testing.T) {
	request := &http.Request{
		Header: http.Header{},
	}

	setHeaders(request)

	assert.NotEqual(t, "", request.Header.Get("Content-Type"))
}

func TestPrepareForm(t *testing.T) {
	request := requests.ConduitConnectRequest{
		Client:        "gonduit",
		ClientVersion: "1",
	}
	marshalled, _ := json.Marshal(request)

	options := &ClientOptions{}

	form, err := prepareForm(&request, options)

	assert.Nil(t, err)
	assert.Equal(t, "true", form.Get("__conduit__"))
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, string(marshalled), form.Get("params"))
}

func TestPrepareForm_withNoBody(t *testing.T) {
	form, err := prepareForm(nil, &ClientOptions{})

	assert.Nil(t, err)
	assert.Equal(t, "", form.Get("__conduit__"))
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, "{}", form.Get("params"))
}

func TestPrepareForm_withAPIToken(t *testing.T) {

	form, err := prepareForm(&requests.Request{}, &ClientOptions{
		APIToken: "hello-world-hello-world",
	})

	jsonBody := `{"__conduit__":{"token":"hello-world-hello-world"}}`

	assert.Nil(t, err)
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, jsonBody, form.Get("params"))
}

func TestPrepareForm_withSessionKey(t *testing.T) {
	form, err := prepareForm(&requests.Request{}, &ClientOptions{
		SessionKey: "hello-world-hello-world",
	})

	jsonBody := `{"__conduit__":{"sessionKey":"hello-world-hello-world"}}`

	assert.Nil(t, err)
	assert.Equal(t, "json", form.Get("output"))
	assert.Equal(t, string(jsonBody), form.Get("params"))
}

func TestPrepareForm_withError(t *testing.T) {
	_, err := prepareForm(
		http.Request{},
		&ClientOptions{},
	)

	assert.NotNil(t, err)
}

func TestHandleConnectRequest(t *testing.T) {
	request := requests.ConduitConnectRequest{}
	request2 := requests.DifferentialQueryRequest{}

	form := url.Values{}
	form2 := url.Values{}

	handleConnectRequest(&form, &request)
	handleConnectRequest(&form2, &request2)

	assert.Equal(t, "true", form.Get("__conduit__"))
	assert.Equal(t, "", form2.Get("__conduit__"))
}
