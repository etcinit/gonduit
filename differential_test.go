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

func TestDifferentialGetCommitPaths(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	response := server.ResponseFromJSON(`{
		"result": [
			"differential.go",
			"differential_test.go"
		]
	}`)
	s.RegisterMethod(DifferentialGetCommitPathsMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitPathsRequest{RevisionID: 123}
	resp, err := c.DifferentialGetCommitPaths(req)
	assert.NoError(t, err)
	want := &responses.DifferentialGetCommitPathsResponse{
		"differential.go",
		"differential_test.go",
	}
	assert.Equal(t, want, resp)
}

func TestDifferentialGetCommitMessage(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	response := server.ResponseFromJSON(`{
		"result": "Commit description."
	}`)
	s.RegisterMethod(DifferentialGetCommitMessageMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitMessageRequest{
		RevisionID: 123,
	}
	resp, err := c.DifferentialGetCommitMessage(req)
	assert.NoError(t, err)
	want := responses.DifferentialGetCommitMessageResponse("Commit description.")
	assert.Equal(t, &want, resp)
}

func TestDifferentialQuery(t *testing.T) {

	tests := map[string]struct {
		apiResponse string
		want        entities.DifferentialRevision
	}{
		"response_with_reviewers": {
			apiResponse: `{
				"result": [{
					"id": "123",
					"reviewers": {
						"PHID-USER-123": "PHID-USER-123"
					}
				}]
			}`,
			want: entities.DifferentialRevision{
				ID: "123",
				Reviewers: entities.DifferentialRevisionReviewers{
					"PHID-USER-123": "PHID-USER-123",
				},
			},
		},
		// Phabricator returns empty slice instead of empty map when reviewers
		// do not exist. And a map when they do. This case should be handled
		// separately when unmarshaling JSON in Golang.
		"response_with_no_reviewers": {
			apiResponse: `{
				"result": [{
					"id": "123",
					"reviewers": []
				}]
			}`,
			want: entities.DifferentialRevision{
				ID: "123",
			},
		},
	}

	req := requests.DifferentialQueryRequest{
		IDs: []uint64{123},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := server.New()
			defer s.Close()
			s.RegisterCapabilities()

			s.RegisterMethod(
				DifferentialQueryMethod, http.StatusOK, server.ResponseFromJSON(test.apiResponse))
			c, err := Dial(s.GetURL(), &core.ClientOptions{
				APIToken: "some-token",
			})
			assert.Nil(t, err)
			resp, err := c.DifferentialQuery(req)
			assert.NoError(t, err)
			want := &responses.DifferentialQueryResponse{
				&test.want,
			}
			assert.Equal(t, want, resp)
		})
	}

}

const differentialRevisionSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 123,
        "type": "DREV",
        "phid": "PHID-DREV-000",
        "fields": {
          "title": "Revision title",
          "uri": "https://secure.phabricator.com/D1",
          "authorPHID": "PHID-USER-000",
          "repositoryPHID": "PHID-REPO-000",
          "diffPHID": "PHID-DIFF-000",
          "summary": "Revision summary",
          "testPlan": "Test plan",
          "isDraft": true,
          "holdAsDraft": true,
          "status": {
            "value": "needs-review",
            "name": "Needs Review",
            "closed": false
          }
        },
        "attachments": {
          "reviewers": {
            "reviewers": [
              {
                "reviewerPHID": "PHID-USER-123",
                "status": "added",
                "isBlocking": true
              }
            ]
          },
          "subscribers": {
            "subscriberPHIDs": [
              "PHID-USER-456"
            ],
            "subscriberCount": 1,
            "viewerIsSubscribed": true
          },
          "projects": {
            "projectPHIDs": [
              "PHID-PROJ-123"
            ]
          }
        }
      }
    ],
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

func TestDifferentialRevisionSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(differentialRevisionSearchResponseJSON)
	s.RegisterMethod(DifferentialRevisionSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialRevisionSearchRequest{
		Constraints: &requests.DifferentialRevisionSearchConstraints{
			IDs: []int{123},
		},
		Attachments: &requests.DifferentialRevisionSearchAttachments{
			Reviewers:   true,
			Subscribers: true,
			Projects:    true,
		},
	}
	resp, err := c.DifferentialRevisionSearch(req)
	assert.NoError(t, err)
	want := responses.DifferentialRevisionSearchResponse{
		Data: []*responses.DifferentialRevisionSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   123,
					Type: "DREV",
					PHID: "PHID-DREV-000",
				},
				Fields: responses.DifferentialRevisionSearchResponseItemFields{
					Title:          "Revision title",
					URI:            "https://secure.phabricator.com/D1",
					AuthorPHID:     "PHID-USER-000",
					RepositoryPHID: "PHID-REPO-000",
					DiffPHID:       "PHID-DIFF-000",
					Summary:        "Revision summary",
					TestPlan:       "Test plan",
					IsDraft:        true,
					HoldAsDraft:    true,
					Status: responses.DifferentialRevisionStatus{
						Value:  "needs-review",
						Name:   "Needs Review",
						Closed: false,
					},
				},
				Attachments: responses.DifferentialRevisionSearchAttachments{
					Reviewers: responses.SearchAttachmentReviewers{
						Reviewers: []responses.AttachmentReviewer{
							{
								ReviewerPHID: "PHID-USER-123",
								Status:       "added",
								IsBlocking:   true,
							},
						},
					},
					Subscribers: responses.SearchAttachmentSubscribers{
						SubscriberPHIDs:    []string{"PHID-USER-456"},
						SubscriberCount:    1,
						ViewerIsSubscribed: true,
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

const differentialDiffSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 123,
        "type": "DIFF",
        "phid": "PHID-DIFF-123",
        "fields": {
          "revisionPHID": "PHID-DREV-123",
          "authorPHID": "PHID-USER-123",
          "repositoryPHID": "PHID-REPO-123",
          "refs": [
            {
              "type": "branch",
              "name": "patch"
            },
            {
              "type": "onto",
              "name": "master"
            },
            {
              "type": "base",
              "identifier": "aaaa"
            }
          ],
          "dateCreated": 1606253067,
          "dateModified": 1606253237,
          "policy": {
            "view": "users"
          }
        },
        "attachments": {
          "commits": {
            "commits": [
              {
                "identifier": "bbb",
                "tree": "ccc",
                "parents": [
                  "ddd"
                ],
                "author": {
                  "name": "Author Subauthor",
                  "email": "author@example.com",
                  "raw": "\"Author Subauthor\" <author@example.com>",
                  "epoch": 1606252972
                },
                "message": "Title\n\nSummary: Summary\n\n"
              }
            ]
          }
        }
      }
    ],
    "maps": [],
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

func TestDifferentialDiffSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(differentialDiffSearchResponseJSON)
	s.RegisterMethod(DifferentialDiffSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialDiffSearchRequest{
		Constraints: &requests.DifferentialDiffSearchConstraints{
			IDs: []int{123},
		},
		Attachments: &requests.DifferentialDiffSearchAttachments{
			Commits: true,
		},
	}
	resp, err := c.DifferentialDiffSearch(req)
	assert.NoError(t, err)
	want := responses.DifferentialDiffSearchResponse{
		Data: []*responses.DifferentialDiffSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   123,
					Type: "DIFF",
					PHID: "PHID-DIFF-123",
				},
				Fields: responses.DifferentialDiffSearchResponseItemFields{
					RevisionPHID:   "PHID-DREV-123",
					AuthorPHID:     "PHID-USER-123",
					RepositoryPHID: "PHID-REPO-123",
					Refs: []responses.DifferentialDiffRef{
						{
							Type: "branch",
							Name: "patch",
						},
						{
							Type: "onto",
							Name: "master",
						},
						{
							Type:       "base",
							Identifier: "aaaa",
						},
					},
					DateCreated:  timestamp(1606253067),
					DateModified: timestamp(1606253237),
				},
				Attachments: responses.DifferentialDiffSearchAttachments{
					Commits: responses.SearchAttachmentCommits{
						Commits: []responses.AttachmentCommit{
							{
								Identifier: "bbb",
								Tree:       "ccc",
								Parents: []string{
									"ddd",
								},
								Author: responses.AttachmentCommitAuthor{
									Name:  "Author Subauthor",
									Email: "author@example.com",
									Raw:   "\"Author Subauthor\" <author@example.com>",
									Epoch: timestamp(1606252972),
								},
								Message: "Title\n\nSummary: Summary\n\n",
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
