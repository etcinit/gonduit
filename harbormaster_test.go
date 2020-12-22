package gonduit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

const buildableSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 54057,
        "type": "HMBB",
        "phid": "PHID-HMBB-6tceawkrkt55btokp7es",
        "fields": {
          "objectPHID": "PHID-DIFF-7grzvaqb24vaorwgj6f6",
          "containerPHID": "PHID-DREV-ea4xglpfvktonm7cyzmq",
          "buildableStatus": {
            "value": "failed"
          },
          "isManual": false,
          "uri": "https://www.example.com/B54057",
          "dateCreated": 1419993553,
          "dateModified": 1419994281,
          "policy": {
            "view": "users",
            "edit": "users"
          }
        },
        "attachments": {}
      }
    ],
    "maps": {},
    "query": {
      "queryKey": null
    },
    "cursor": {
      "limit": 100,
      "after": null,
      "before": null,
      "order": null
    }
  }
}`

func TestHarbormasterBuildableSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(buildableSearchResponseJSON)
	s.RegisterMethod(HarbormasterBuildableSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.HarbormasterBuildableSearchRequest{
		Constraints: &requests.HarbormasterBuildableSearchConstraints{
			IDs: []int{54057},
		},
	}
	resp, err := c.HarbormasterBuildableSearch(req)
	assert.NoError(t, err)
	want := responses.HarbormasterBuildableSearchResponse{
		Data: []*responses.HarbormasterBuildableSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   54057,
					Type: "HMBB",
					PHID: "PHID-HMBB-6tceawkrkt55btokp7es",
				},
				Fields: responses.HarbormasterBuildableSearchResponseItemFields{
					ObjectPHID:    "PHID-DIFF-7grzvaqb24vaorwgj6f6",
					ContainerPHID: "PHID-DREV-ea4xglpfvktonm7cyzmq",
					BuildableStatus: responses.BuildableStatus{
						Value: "failed",
					},
					IsManual:     false,
					URI:          "https://www.example.com/B54057",
					DateCreated:  timestamp(1419993553),
					DateModified: timestamp(1419994281),
				},
			},
		},
		Cursor: responses.SearchCursor{
			Limit: 100,
		},
	}
	assert.Equal(t, &want, resp)
}
