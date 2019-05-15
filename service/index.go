package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Index of UB
func (ub *UrantiaBook) Index(ctx context.Context, req *api.IndexRequest) (*api.IndexResponse, error) {
	var resp = &api.IndexResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	size := req.Size
	if req.Size <= 0 || req.Size > len(UBPapers) {
		size = len(UBPapers)
	}
	skip := req.Skip
	if req.Skip <= 0 || req.Skip > len(UBPapers) {
		skip = 0
	}

	for i := skip; i < size+skip && i < len(UBPapers); i++ {
		p := UBPapers[i]
		paper := api.PaperIndex{ID: p.ID, Title: p.Title, Author: p.Author}
		for _, s := range p.Sections {
			paper.Sections = append(paper.Sections, api.SectionIndex{ID: s.ID, Title: s.Title})
		}
		resp.Papers = append(resp.Papers, paper)
	}
	resp.Total = uint64(len(UBPapers))

	return resp, nil
}

// MakeIndexEndpoint -
func MakeIndexEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.IndexRequest)
		return svc.Index(ctx, &req)
	}
}

// DecodeIndexRequest -
func DecodeIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.IndexRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
