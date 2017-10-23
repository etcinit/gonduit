package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// ProjectQuery performs a call to project.query.
func (c *Conn) ProjectQuery(
	req requests.ProjectQueryRequest,
) (*responses.ProjectQueryResponse, error) {
	var res responses.ProjectQueryResponse

	if err := c.Call("project.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectCreate performs a call to project.create.
func (c *Conn) ProjectCreate(
	req requests.ProjectCreateRequest,
) (*responses.ProjectCreateResponse, error) {
	var res responses.ProjectCreateResponse

	if err := c.Call("project.create", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
