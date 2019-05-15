package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// List Pages
func (Page) List(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	var resp = api.ListResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	resp.Parts = UB.Parts

	for i := req.Skip; i <= req.Size; i++ {
		resp.Papers = append(resp.Papers, UB.Papers[i])
	}
	resp.Total = uint64(len(UB.Papers))

	return &resp, nil
}

// MakeListEndpoint -
func MakeListEndpoint(svc Page) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ListRequest)
		return svc.List(ctx, &req)
	}
}

// DecodeListRequest -
func DecodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
