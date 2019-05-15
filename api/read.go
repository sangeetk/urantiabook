package api

// ReadRequest for the UB paper
type ReadRequest struct {
	Language string `json:"language"`
	Slug     string `json:"slug"`
}

// ReadResponse for the UB paper
type ReadResponse struct {
	Language string       `json:"language"`
	Request  *ReadRequest `json:"request"`
	Section  Section      `json:"section"`
	Err      string       `json:"err"`
}
