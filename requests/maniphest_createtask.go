package requests

// ManiphestCreateTaskRequest represents a request to maniphest.createtask.
type ManiphestCreateTaskRequest struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	OwnerPHID    string   `json:"ownerPHID,omitempty"`
	ViewPolicy   string   `json:"viewPolicy,omitempty"`
	EditPolicy   string   `json:"editPolicy,omitempty"`
	CCPHIDs      []string `json:"ccPHIDs,omitempty"`
	Priority     int      `json:"priority,omitempty"`
	ProjectPHIDs []string `json:"projectPHIDs,omitempty"`
	Request
}
