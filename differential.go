package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// DifferentialQuery performs a call to differential.query.
func (c *Conn) DifferentialQuery(
	req requests.DifferentialQueryRequest,
) (*responses.DifferentialQueryResponse, error) {
	var res responses.DifferentialQueryResponse

	if err := c.Call("differential.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialQuery performs a call to differential.querydiffs.
func (c *Conn) DifferentialQueryDiffs(
	req requests.DifferentialQueryDiffsRequest,
) (*responses.DifferentialQueryDiffsResponse, error) {
	var res responses.DifferentialQueryDiffsResponse

	if err := c.Call("differential.querydiffs", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
