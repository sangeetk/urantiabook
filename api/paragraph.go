package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ParagraphRequest for the UB paper
type ParagraphRequest struct {
	Language string `json:"language"`
	ID       string `json:"id"`
}

type Paragraph struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// ParagraphResponse for the UB paper
type ParagraphResponse struct {
	Language  string            `json:"language"`
	Request   *ParagraphRequest `json:"request"`
	Paragraph *Paragraph        `json:"paragraph"`
	Err       string            `json:"err,omitempty"`
}

// GetParagraph from a seciton of UB Paper
func GetParagraph(req *ParagraphRequest, dns string) (*ParagraphResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/paragraph")
	if err != nil {
		log.Fatal(err.Error())
	}

	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeParagraphResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(ParagraphResponse).Err != "" {
		return nil, errors.New(resp.(ParagraphResponse).Err)
	}

	response := resp.(ParagraphResponse)
	return &response, nil
}

// decodeParagraphResponse decodes the paragraph
func decodeParagraphResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ParagraphResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
