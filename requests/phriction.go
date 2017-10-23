package requests

// PhrictionInfoRequest represents a request to phriction.info API endpoint
type PhrictionInfoRequest struct {
	Slug string `json:"slug"`
	Request
}

// PhrictionCreateRequest represents a request to phriction.create API endpoint
type PhrictionCreateRequest struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Decription string `json:"description,omitempty"`
	Request
}
