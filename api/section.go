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

// SectionRequest for the UB paper
type SectionRequest struct {
	Language string `json:"language"`
	ID       string `json:"id"`
}

// Section of a UB paper
type Section struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// SectionResponse for the UB paper
type SectionResponse struct {
	Language string          `json:"language"`
	Request  *SectionRequest `json:"request"`
	Section  *Section        `json:"section"`
	Err      string          `json:"err,omitempty"`
}

// GetSection of a UB Paper
func GetSection(req *SectionRequest, dns string) (*SectionResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/section")
	if err != nil {
		log.Fatal(err.Error())
	}

	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeSectionResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(SectionResponse).Err != "" {
		return nil, errors.New(resp.(SectionResponse).Err)
	}

	response := resp.(SectionResponse)
	return &response, nil
}

// decodeSectionResponse decodes the search results
func decodeSectionResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response SectionResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
