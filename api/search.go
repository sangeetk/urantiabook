package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	ht "github.com/urantiatech/kit/transport/http"
)

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
	Hits     []interface{}  `json:"hits"`
	Total    uint64         `json:"total_hits"`
	Took     time.Duration  `json:"took"`
	Err      string         `json:"err,omitempty"`
}

// Search Pages
func Search(req *SearchRequest, dns string) (*SearchResults, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/search")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeSearchResults).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(SearchResults).Err != "" {
		return nil, errors.New(resp.(SearchResults).Err)
	}
	response := resp.(SearchResults)
	return &response, nil
}

// decodeSearchResults decodes the search results
func decodeSearchResults(ctx context.Context, r *http.Response) (interface{}, error) {
	var response SearchResults
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
