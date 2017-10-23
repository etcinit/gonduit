package requests

import "github.com/etcinit/gonduit/constants"

// ProjectQueryRequest represents a request to project.query.
type ProjectQueryRequest struct {
	IDs     []string                `json:"ids"`
	Names   []string                `json:"names"`
	PHIDs   []string                `json:"phids"`
	Slugs   []string                `json:"slugs"`
	Icons   []string                `json:"icons"`
	Colors  []string                `json:"colors"`
	Status  constants.ProjectStatus `json:"status"`
	Members []string                `json:"members"`
	Limit   uint64                  `json:"limit"`
	Offset  uint64                  `json:"offset"`
	Request
}

// ProjectCreateRequest represents a request to project.create.
type ProjectCreateRequest struct {
	Name          string   `json:"name"`
	MemberPHIDs   []string `json:"member_phids,omitempty"`
	MemberAliases []string `json:"member_aliases,omitempty"`
	Request
}
