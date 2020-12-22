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

const transactionSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 123,
        "phid": "PHID-XACT-DREV-123",
        "type": "close",
        "authorPHID": "PHID-USER-123",
        "objectPHID": "PHID-DREV-123",
        "dateCreated": 1606741970,
        "dateModified": 1606741970,
        "groupID": "123456",
        "comments": [
          {
            "id": 38153297,
            "phid": "PHID-XCMT-123",
            "version": 1,
            "authorPHID": "PHID-USER-123",
            "dateCreated": 1606679165,
            "dateModified": 1606679165,
            "removed": false,
            "content": {
              "raw": "some comment"
            }
          }
        ],
        "fields": {
          "commitPHIDs": [
            "PHID-CMIT-l4cgf7tpkwvq45zo4sd6"
          ],
          "old": "PHID-DIFF-t3vde55jeilu47spbnve",
          "new": "PHID-DIFF-qtiogk4axoi4kytpaken",
          "operations": [
            {
              "operation": "add",
              "phid": "PHID-PROJ-a3xyvhtxaenwd6hqslzg",
              "oldStatus": null,
              "newStatus": "added",
              "isBlocking": false
            }
          ]
        }
      },
      {
        "id": 1234,
        "phid": "PHID-XACT-DREV-1234",
        "type": null,
        "authorPHID": "PHID-USER-1234",
        "objectPHID": "PHID-DREV-123",
        "dateCreated": 1606741970,
        "dateModified": 1606741970,
        "groupID": "123456",
        "comments": [],
        "fields": []
      }
    ],
    "cursor": {
      "limit": 100,
      "after": null,
      "before": null
    }
  }
}`

func TestTransactionSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(transactionSearchResponseJSON)
	s.RegisterMethod(TransactionSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.TransactionSearchRequest{
		ObjectIdentifier: "D123",
	}
	resp, err := c.TransactionSearch(req)
	want := responses.TransactionSearchResponse{
		Data: []*responses.TransactionSearchResponseItem{
			&responses.TransactionSearchResponseItem{
				ResponseObject: responses.ResponseObject{
					ID:   123,
					Type: "close",
					PHID: "PHID-XACT-DREV-123",
				},
				Fields: responses.TransactionSearchResponseItemFields{
					Old: "PHID-DIFF-t3vde55jeilu47spbnve",
					New: "PHID-DIFF-qtiogk4axoi4kytpaken",
					Operations: []responses.TransactionSearchResponseItemFieldsOperation{
						{
							Operation:  "add",
							PHID:       "PHID-PROJ-a3xyvhtxaenwd6hqslzg",
							OldStatus:  "",
							NewStatus:  "added",
							IsBlocking: false,
						},
					},
					CommitPHIDs: []string{
						"PHID-CMIT-l4cgf7tpkwvq45zo4sd6",
					},
				},
				AuthorPHID:   "PHID-USER-123",
				ObjectPHID:   "PHID-DREV-123",
				GroupID:      "123456",
				DateCreated:  timestamp(1606741970),
				DateModified: timestamp(1606741970),
				Comments: []responses.TransactionSearchResponseItemComment{
					{
						ID:           38153297,
						PHID:         "PHID-XCMT-123",
						Version:      1,
						AuthorPHID:   "PHID-USER-123",
						DateCreated:  timestamp(1606679165),
						DateModified: timestamp(1606679165),
						Removed:      false,
						Content: responses.TransactionSearchResponseItemContent{
							Raw: "some comment",
						},
					},
				},
			},
			&responses.TransactionSearchResponseItem{
				ResponseObject: responses.ResponseObject{
					ID:   1234,
					PHID: "PHID-XACT-DREV-1234",
				},
				AuthorPHID:   "PHID-USER-1234",
				ObjectPHID:   "PHID-DREV-123",
				GroupID:      "123456",
				DateCreated:  timestamp(1606741970),
				DateModified: timestamp(1606741970),
				Comments:     []responses.TransactionSearchResponseItemComment{},
			},
		},
		Cursor: entities.Cursor{
			Limit: 100,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, &want, resp)
}
