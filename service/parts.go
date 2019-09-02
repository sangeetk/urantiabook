package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/text/language"
)

// Parts of UB
func (ub *UrantiaBook) Parts(ctx context.Context, req *api.PartsRequest) (*api.PartsResponse, error) {
	var resp = &api.PartsResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	for _, p := range UBParts {
		part := api.Part{
			ID:         p.ID,
			Title:      p.Title,
			PaperStart: p.PaperStart,
			PaperEnd:   p.PaperEnd,
			Authors:    p.Authors,
		}
		resp.Parts = append(resp.Parts, part)
	}

	return resp, nil
}

// MakePartsEndpoint -
func MakePartsEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.PartsRequest)
		return svc.Parts(ctx, &req)
	}
}

// DecodePartsRequest -
func DecodePartsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.PartsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
