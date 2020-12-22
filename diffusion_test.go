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

const repositorySearchResponseJSON = `{
  "result": {
  "data": [
    {
      "id": 1000,
      "type": "REPO",
      "phid": "PHID-REPO-1000",
      "fields": {
        "name": "Gonduit",
        "vcs": "git",
        "callsign": "GND",
        "shortName": null,
        "status": "active",
        "isImporting": false,
        "defaultBranch": "master",
        "description": {
          "raw": "description"
        },
        "spacePHID": "PHID-SPCE-1000",
        "dateCreated": 1453225096,
        "dateModified": 1604986358
      },
      "attachments": {
        "uris": {
          "uris": [
            {
              "id": "34552",
              "type": "RURI",
              "phid": "PHID-RURI-ztifpwbrbha7rfbbk6ai",
              "fields": {
                "repositoryPHID": "PHID-REPO-meb4ivps5qj5gtlkfc7v",
                "uri": {
                  "raw": "git@github.com:uber/gonduit.git"
                },
                "credentialPHID": "PHID-CDTL-33r3eatdjwks355etw47",
                "disabled": false,
                "dateCreated": "1489737532",
                "dateModified": "1489737532"
              }
            }
          ]
        },
        "metrics": {
          "commitCount": 721
        },
        "projects": {
          "projectPHIDs": [
            "PHID-PROJ-123"
          ]
        }
      }
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

func TestDiffusionRepositorySearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(repositorySearchResponseJSON)
	s.RegisterMethod(DiffusionRepositorySearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DiffusionRepositorySearchRequest{
		Constraints: &requests.DiffusionRepositorySearchConstraints{
			IDs: []int{1000},
		},
	}
	resp, err := c.DiffusionRepositorySearch(req)
	assert.NoError(t, err)
	want := responses.DiffusionRepositorySearchResponse{
		Data: []*responses.DiffusionRepositorySearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   1000,
					Type: "REPO",
					PHID: "PHID-REPO-1000",
				},
				Fields: responses.DiffusionRepositorySearchResponseItemFields{
					Name:          "Gonduit",
					VCS:           "git",
					Callsign:      "GND",
					ShortName:     "",
					Status:        "active",
					IsImporting:   false,
					DefaultBranch: "master",
					Description: responses.DiffusionRepositoryDescription{
						Raw: "description",
					},
					SpacePHID:    "PHID-SPCE-1000",
					DateCreated:  timestamp(1453225096),
					DateModified: timestamp(1604986358),
				},
				Attachments: responses.DiffusionRepositorySearchAttachments{
					URIs: responses.SearchAttachmentURIs{
						URIs: []responses.RepositoryURIItem{
							{
								Fields: responses.RepositoryURIItemFields{
									URI: responses.RepositoryURI{
										Raw: "git@github.com:uber/gonduit.git",
									},
									Disabled:     false,
									DateCreated:  timestamp(1489737532),
									DateModified: timestamp(1489737532),
								},
							},
						},
					},
					Metrics: responses.SearchAttachmentMetrics{
						CommitCount: 721,
					},
					Projects: responses.SearchAttachmentProjects{
						ProjectPHIDs: []string{"PHID-PROJ-123"},
					},
				},
			},
		},
		Cursor: responses.SearchCursor{
			Limit: 100,
		},
	}
	assert.Equal(t, &want, resp)
}
