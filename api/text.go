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

// TextRequest for the UB paper
type TextRequest struct {
	Language string `json:"language"`
	ID       string `json:"id"`
}

// TextResponse for the UB Paper
type TextResponse struct {
	Language string       `json:"language"`
	Request  *TextRequest `json:"request"`
	Text     string       `json:"text"`
	Err      string       `json:"err,omitempty"`
}

// Text from UB
func Text(req *TextRequest, dns string) (*TextResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/text")
	if err != nil {
		log.Fatal(err.Error())
	}

	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeTextResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(TextResponse).Err != "" {
		return nil, errors.New(resp.(TextResponse).Err)
	}

	response := resp.(TextResponse)
	return &response, nil
}

// decodeTextResponse decodes the response
func decodeTextResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response TextResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
