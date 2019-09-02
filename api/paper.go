package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	ht "github.com/go-kit/kit/transport/http"
)

// PaperRequest for the UB paper
type PaperRequest struct {
	Language string `json:"language"`
	Paper    int    `json:"paper"`
}

// Paragraph within a section
type Paragraph struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	List bool   `json:"list"`
}

// Section of a UB paper
type Section struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	Paragraphs []Paragraph `json:"paragraphs"`
}

// Paper from UB
type Paper struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
	Author   string    `json:"author"`
}

// PaperResponse for the UB Paper
type PaperResponse struct {
	Language string        `json:"language"`
	Request  *PaperRequest `json:"request"`
	Paper    *Paper        `json:"paper"`
	Err      string        `json:"err,omitempty"`
}

// GetPaper from UB
func GetPaper(req *PaperRequest, dns string) (*PaperResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/paper")
	if err != nil {
		log.Fatal(err.Error())
	}

	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodePaperResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(PaperResponse).Err != "" {
		return nil, errors.New(resp.(PaperResponse).Err)
	}

	response := resp.(PaperResponse)
	return &response, nil
}

// decodePaperResponse decodes the response
func decodePaperResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response PaperResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
