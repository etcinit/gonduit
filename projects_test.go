package gonduit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

const projectSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 10743,
        "type": "PROJ",
        "phid": "PHID-PROJ-project",
        "fields": {
          "name": "IMPORTANT",
          "slug": "IMP",
          "subtype": "default",
          "milestone": 26,
          "depth": 0,
          "parent": {
            "id": 10742,
            "phid": "PHID-PROJ-parent",
            "name": "PAR"
          },
          "icon": {
            "key": "project",
            "name": "Project",
            "icon": "fa-briefcase"
          },
          "color": {
            "key": "blue",
            "name": "Blue"
          },
          "spacePHID": null,
          "dateCreated": 1483575604,
          "dateModified": 1607595030,
          "policy": {
            "view": "PHID-PROJ-prof",
            "edit": "PHID-PROJ-prof",
            "join": "PHID-PROJ-prof"
          },
          "description": "Description"
        },
        "attachments": {
          "members": {
            "members": [
              {
                "phid": "PHID-USER-some"
              }
            ]
          },
          "watchers": {
            "watchers": [
              {
                "phid": "PHID-USER-some2"
              }
            ]
          },
          "ancestors": {
            "ancestors": [
              {
                "id": 10742,
                "phid": "PHID-PROJ-parent",
                "name": "PAR"
              }
            ]
          }
        }
      }
    ],
    "maps": {
      "slugMap": {}
    },
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

func TestProjectSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(projectSearchResponseJSON)
	s.RegisterMethod(ProjectSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.ProjectSearchRequest{
		Constraints: &requests.ProjectSearchConstraints{
			IDs: []int{123},
		},
		Attachments: &requests.ProjectSearchAttachments{
			Ancestors: true,
			Members:   true,
			Watchers:  true,
		},
	}
	resp, err := c.ProjectSearch(req)
	require.NoError(t, err)
	want := responses.ProjectSearchResponse{
		Data: []*responses.ProjectSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   10743,
					Type: "PROJ",
					PHID: "PHID-PROJ-project",
				},
				Fields: responses.ProjectSearchResponseItemFields{
					Name:        "IMPORTANT",
					Slug:        "IMP",
					Description: "Description",
					Subtype:     "default",
					Milestone:   26,
					Parent: &responses.ProjectParent{
						ID:   10742,
						PHID: "PHID-PROJ-parent",
						Name: "PAR",
					},
					Icon: responses.ProjectIcon{
						Key:  "project",
						Name: "Project",
						Icon: "fa-briefcase",
					},
					Color: responses.ProjectColor{
						Key:  "blue",
						Name: "Blue",
					},
					DateCreated:  timestamp(1483575604),
					DateModified: timestamp(1607595030),
				},
				Attachments: responses.ProjectSearchAttachments{
					Members: responses.SearchAttachmentMembers{
						Members: []responses.AttachmentMember{
							{
								PHID: "PHID-USER-some",
							},
						},
					},
					Watchers: responses.SearchAttachmentWatchers{
						Watchers: []responses.AttachmentWatcher{
							{
								PHID: "PHID-USER-some2",
							},
						},
					},
					Ancestors: responses.SearchAttachmentAncestors{
						Ancestors: []responses.ProjectParent{
							{
								ID:   10742,
								PHID: "PHID-PROJ-parent",
								Name: "PAR",
							},
						},
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
