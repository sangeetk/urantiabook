package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Paper from UB
func (ub *UrantiaBook) Paper(ctx context.Context, req *api.PaperRequest) (*api.PaperResponse, error) {
	var resp = &api.PaperResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil || id < 0 || int(id) >= len(UBPapers) {
		resp.Err = "Invalid Request"
		return resp, nil
	}

	resp.Paper = &UBPapers[id]

	return resp, nil
}

// MakePaperEndpoint -
func MakePaperEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.PaperRequest)
		return svc.Paper(ctx, &req)
	}
}

// DecodePaperRequest -
func DecodePaperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.PaperRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
