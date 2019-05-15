package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/blevesearch/bleve"
	q "github.com/blevesearch/bleve/search/query"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Search UrantiaBook
func (ub *UrantiaBook) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResults, error) {
	var resp = &api.SearchResults{Request: req}
	var searchRequest *bleve.SearchRequest
	var query q.Query

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	if req.Query == "" {
		query = bleve.NewMatchAllQuery()
	} else if req.Fuzzy {
		query = bleve.NewFuzzyQuery(req.Query)
	} else {
		query = bleve.NewQueryStringQuery(req.Query)
	}

	// Create a new search request
	searchRequest = bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchRequest.Highlight = bleve.NewHighlight()
	searchRequest.Size = req.Size
	if searchRequest.Size <= 0 {
		searchRequest.Size = 10
	}
	searchRequest.From = req.Skip
	searchResult, err := Index.Search(searchRequest)
	if err != nil {
		resp.Err = "Not Found"
		return resp, nil
	}

	resp.Total = searchResult.Total
	resp.Took = searchResult.Took

	for _, hit := range searchResult.Hits {
		resp.Hits = append(resp.Hits, hit.Fields)
	}

	return resp, nil
}

// MakeSearchEndpoint -
func MakeSearchEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SearchRequest)
		return svc.Search(ctx, &req)
	}
}

// DecodeSearchRequest -
func DecodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
