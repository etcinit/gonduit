package gonduit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/entities"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

const edgeSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "sourcePHID": "PHID-TASK-123",
        "edgeType": "mention",
        "destinationPHID": "PHID-TASK-100"
      },
      {
        "sourcePHID": "PHID-TASK-123",
        "edgeType": "mention",
        "destinationPHID": "PHID-DREV-200"
      }
    ],
    "cursor": {
      "limit": 100,
      "after": null,
      "before": null
    }
  }
}`

func TestEdgeSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(edgeSearchResponseJSON)
	s.RegisterMethod(EdgeSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.EdgeSearchRequest{
		SourcePHIDs: []string{"PHID-TASK-123"},
		Types: []entities.EdgeType{
			entities.EdgeMention,
		},
	}
	resp, err := c.EdgeSearch(req)
	want := responses.EdgeSearchResponse{
		Data: []entities.Edge{
			{
				SourcePHID:      "PHID-TASK-123",
				DestinationPHID: "PHID-TASK-100",
				EdgeType:        entities.EdgeMention,
			},
			{
				SourcePHID:      "PHID-TASK-123",
				DestinationPHID: "PHID-DREV-200",
				EdgeType:        entities.EdgeMention,
			},
		},

		Cursor: entities.Cursor{
			Limit: 100,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, &want, resp)
}
