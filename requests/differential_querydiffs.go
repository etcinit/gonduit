package requests

// DifferentialQueryRequest represents a request
// to the differential.querydiffs call.
type DifferentialQueryDiffsRequest struct {
	IDs         []uint64 `json:"ids"`
	RevisionIDs []uint64 `json:"revisionIDs"`
	Request
}
