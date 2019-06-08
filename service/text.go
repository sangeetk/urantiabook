package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/text/language"
)

// Text from UB
func (ub *UrantiaBook) Text(ctx context.Context, req *api.TextRequest) (*api.TextResponse, error) {
	var resp = &api.TextResponse{Request: req}

	// Set Language code
	if req.Language == "" {
		req.Language = language.English.String()
	}
	resp.Language = req.Language

	ids := strings.Split(req.ID, ":")
	if len(ids) != 2 {
		resp.Err = "Invalid Request"
		log.Println("ids", ids)
		return resp, nil
	}

	// Get Paper ID
	paperid, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil || paperid < 0 || int(paperid) >= len(UBPapers) {
		resp.Err = "Paper Not Found"
		return resp, nil
	}

	idds := strings.Split(ids[1], ".")
	if len(idds) != 2 {
		resp.Err = "Invalid Section and Paragraph IDs"
		return resp, nil
	}

	// Get Section ID
	sectionid, err := strconv.ParseInt(idds[0], 10, 64)
	if err != nil || sectionid < 0 || int(sectionid) >= len(UBPapers[paperid].Sections) {
		resp.Err = "Section Not Found"
		return resp, nil
	}

	// Get Paragraph ID
	paraid, err := strconv.ParseInt(idds[1], 10, 64)
	if err != nil || paraid <= 0 || int(paraid) > len(UBPapers[paperid].Sections[sectionid].Paragraphs) {
		resp.Err = "Paragraph Not Found"
		return resp, nil
	}

	resp.Text = UBPapers[paperid].Sections[sectionid].Paragraphs[paraid-1].Text
	return resp, nil
}

// MakeTextEndpoint -
func MakeTextEndpoint(svc UrantiaBook) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.TextRequest)
		return svc.Text(ctx, &req)
	}
}

// DecodeTextRequest -
func DecodeTextRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.TextRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
