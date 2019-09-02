package service

/*

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/text/language"
)

// Section from UB Paper
func (ub *UrantiaBook) Section(ctx context.Context, req *api.SectionRequest) (*api.SectionResponse, error) {
	var resp = &api.SectionResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	ids := strings.Split(req.ID, ":")
	if len(ids) != 2 {
		resp.Err = "Not Found"
		return resp, nil
	}

	paperId, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil || paperId < 0 || int(paperId) >= len(UBPapers) {
		resp.Err = "Not Found"
		return resp, nil
	}
	paper := &UBPapers[paperId]

	sectionId, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil || sectionId < 0 || int(sectionId) >= len(paper.Sections) {
		resp.Err = "Not Found"
		return resp, nil
	}

	resp.Section = &paper.Sections[sectionId]

	return resp, nil
}

// MakeSectionEndpoint -
func MakeSectionEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SectionRequest)
		return svc.Section(ctx, &req)
	}
}

// DecodeSectionRequest -
func DecodeSectionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.SectionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
*/
