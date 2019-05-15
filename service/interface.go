package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
)

// UrantiaBookService - interface
type UrantiaBookService interface {
	List(context.Context, *api.ListRequest) (*api.ListResponse, error)
	Read(context.Context, *api.ReadRequest) (*api.ReadResponse, error)
	Search(context.Context, *api.SearchRequest) (*api.SearchResults, error)
}

// Page - Wrapper for UrantiaBookService Interface
type Page struct{}

// EncodeResponse -
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
