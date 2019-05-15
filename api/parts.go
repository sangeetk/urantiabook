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

// PartsRequest
type PartsRequest struct {
	Language string `json:"language"`
}

type Part struct {
	Title      string `json:"title"`
	PaperStart int    `json:"pstart"`
	PaperEnd   int    `json:"pend"`
	Authors    string `json:"authors"`
}

// PartsResponse
type PartsResponse struct {
	Language string        `json:"language"`
	Request  *PartsRequest `json:"request"`
	Parts    []Part        `json:"parts"`
	Err      string        `json:"err,omitempty"`
}

// GetParts of UB papers
func GetParts(req *PartsRequest, dns string) (*PartsResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/parts")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodePartsResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(PartsResponse).Err != "" {
		return nil, errors.New(resp.(PartsResponse).Err)
	}
	response := resp.(PartsResponse)
	return &response, nil
}

// decodePartsResponse decodes the index results
func decodePartsResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response PartsResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
