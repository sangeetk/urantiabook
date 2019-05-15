package api

// SearchRequest - search request
type SearchRequest struct {
	Language string `json:"language"`
	Query    string `json:"query"`
	Fuzzy    bool   `json:"fuzzy"`
	Size     int    `json:"size"`
	Skip     int    `json:"skip"`
}

type Result struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
}

// SearchResults - search results
type SearchResults struct {
	Language string         `json:"language"`
	Request  *SearchRequest `json:"request"`
	Results  []Result       `json:"papers"`
	Total    uint64         `json:"total"`
	Err      string         `json:"err,omitempty"`
}
