package service

/*
import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Search Pages
func (UrantiaBookService) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResults, error) {
	var results = api.SearchResults{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	results.Language = req.Language

	return &results, nil
}

// MakeSearchEndpoint -
func MakeSearchEndpoint(svc Page) endpoint.Endpoint {
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

*/
