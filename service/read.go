package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Read a page from database
func (Page) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {
	var resp = api.ReadResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	return &resp, nil
}

// MakeReadEndpoint -
func MakeReadEndpoint(svc Page) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ReadRequest)
		return svc.Read(ctx, &req)
	}
}

// DecodeReadRequest -
func DecodeReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ReadRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
