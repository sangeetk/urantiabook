package api

// ListRequest - list request
type ListRequest struct {
	Language string `json:"language"`
	Size     int    `json:"size"`
	Skip     int    `json:"skip"`
}

// ListResponse - list results
type ListResponse struct {
	Language string       `json:"language"`
	Request  *ListRequest `json:"request"`
	Parts    []Part       `json:"parts"`
	Papers   []Paper      `json:"papers"`
	Total    uint64       `json:"total"`
	Err      string       `json:"err,omitempty"`
}
