package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
)

// UrantiaBookInterface - interface
type UrantiaBookInterface interface {
	Index(context.Context, *api.IndexRequest) (*api.IndexResponse, error)
	Parts(context.Context, *api.PartsRequest) (*api.PartsResponse, error)
	Paper(context.Context, *api.PaperRequest) (*api.PaperResponse, error)
	Text(context.Context, *api.TextRequest) (*api.TextResponse, error)
	Search(context.Context, *api.SearchRequest) (*api.SearchResults, error)
}

// UrantiaBook - Wrapper for UrantiaBookInterface
type UrantiaBook struct{}

// EncodeResponse -
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
