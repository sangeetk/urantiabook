package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// IndexRequest - index request
type IndexRequest struct {
	Language string `json:"language"`
	Size     int    `json:"size"`
	Skip     int    `json:"skip"`
}

type SectionIndex struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type PaperIndex struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Sections []SectionIndex `json:"sections"`
	Author   string         `json:"author"`
}

// IndexResponse of papers
type IndexResponse struct {
	Language string        `json:"language"`
	Request  *IndexRequest `json:"request"`
	Papers   []PaperIndex  `json:"papers"`
	Total    uint64        `json:"total"`
	Err      string        `json:"err,omitempty"`
}

// GetIndex of UB papers
func GetIndex(req *IndexRequest, dns string) (*IndexResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/index")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeIndexResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(IndexResponse).Err != "" {
		return nil, errors.New(resp.(IndexResponse).Err)
	}
	response := resp.(IndexResponse)
	return &response, nil
}

// encodeRequest encodes the request as JSON
func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeIndexResponse decodes the response
func decodeIndexResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response IndexResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
