package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// PhrictionInfo performs a call to phriction.info
func (c *Conn) PhrictionInfo(
	req requests.PhrictionInfoRequest,
) (*responses.PhrictionInfoResponse, error) {
	var res responses.PhrictionInfoResponse

	if err := c.Call("phriction.info", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// PhrictionCreate performs a call to phriction.create API endpoint.
func (c *Conn) PhrictionCreate(
	req requests.PhrictionCreateRequest,
) (*responses.PhrictionCreateResponse, error) {
	var res responses.PhrictionCreateResponse

	if err := c.Call("phriction.create", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
