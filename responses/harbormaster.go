package responses

import (
	"github.com/uber/gonduit/entities"
	"github.com/uber/gonduit/util"
)

// HarbormasterBuildableSearchResponse contains fields that are in server
// response to differential.revision.search.
type HarbormasterBuildableSearchResponse struct {
	// Data contains search results.
	Data []*HarbormasterBuildableSearchResponseItem `json:"data"`

	// Curson contains paging data.
	Cursor SearchCursor `json:"cursor,omitempty"`
}

// HarbormasterBuildableSearchResponseItem contains information about a
// particular search result.
type HarbormasterBuildableSearchResponseItem struct {
	ResponseObject
	Fields HarbormasterBuildableSearchResponseItemFields `json:"fields"`
	SearchCursor
}

// HarbormasterBuildableSearchResponseItemFields is a collection of object
// fields.
type HarbormasterBuildableSearchResponseItemFields struct {
	ObjectPHID      string             `json:"ObjectPHID"`
	ContainerPHID   string             `json:"ContainerPHID"`
	BuildableStatus BuildableStatus    `json:"buildableStatus"`
	IsManual        bool               `json:"isManual"`
	URI             string             `json:"uri"`
	DateCreated     util.UnixTimestamp `json:"dateCreated"`
	DateModified    util.UnixTimestamp `json:"dateModified"`
}

// BuildableStatus is a container of status value.
type BuildableStatus struct {
	Value entities.BuildableStatus `json:"value"`
}
