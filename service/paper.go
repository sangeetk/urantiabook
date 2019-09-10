package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"git.urantiatech.com/urantiabook/urantiabook/api"
	"github.com/go-kit/kit/endpoint"
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

	if req.Paper < 0 || int(req.Paper) >= len(UBPapers) {
		resp.Err = "Not Found"
		return resp, nil
	}

	if !req.Plaintext {
		resp.Paper = &UBPapers[req.Paper]
		return resp, nil
	}

	// Create a copy of Paper and clean HTML tags
	var paper api.Paper
	paper.ID = UBPapers[req.Paper].ID
	paper.Title = UBPapers[req.Paper].Title
	paper.Author = UBPapers[req.Paper].Author

	for _, sec := range UBPapers[req.Paper].Sections {
		var section = api.Section{ID: sec.ID, Title: sec.Title, Text: sec.Text}

		for _, para := range sec.Paragraphs {
			var paragraph = api.Paragraph{ID: para.ID, Text: para.Text, List: para.List}

			// clean all <em> tags
			paragraph.Text = strings.ReplaceAll(paragraph.Text, "<em>", "")
			paragraph.Text = strings.ReplaceAll(paragraph.Text, "</em>", "")
			section.Paragraphs = append(section.Paragraphs, paragraph)
		}
		section.Text = strings.ReplaceAll(section.Text, "<em>", "")
		section.Text = strings.ReplaceAll(section.Text, "</em>", "")
		paper.Sections = append(paper.Sections, section)
	}

	resp.Paper = &paper

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
