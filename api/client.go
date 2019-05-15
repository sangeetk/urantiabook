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

// List Pages
func List(req *ListRequest, dns string) (*ListResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/list")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeListResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(ListResponse).Err != "" {
		return nil, errors.New(resp.(ListResponse).Err)
	}
	response := resp.(ListResponse)
	return &response, nil
}

// Read a page from database
func Read(req *ReadRequest, dns string) (*ReadResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/read")
	if err != nil {
		log.Fatal(err.Error())
	}

	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeReadResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(ReadResponse).Err != "" {
		return nil, errors.New(resp.(ReadResponse).Err)
	}

	response := resp.(ReadResponse)
	return &response, nil
}

// Search Pages
func Search(req *SearchRequest, dns string) (*SearchResults, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/search")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeSearchResults).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.(SearchResults).Err != "" {
		return nil, errors.New(resp.(SearchResults).Err)
	}
	response := resp.(SearchResults)
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

// decodeListResponse decodes the list results
func decodeListResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ListResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// decodeReadResponse decodes the search results
func decodeReadResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ReadResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// decodeSearchResults decodes the search results
func decodeSearchResults(ctx context.Context, r *http.Response) (interface{}, error) {
	var response SearchResults
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
