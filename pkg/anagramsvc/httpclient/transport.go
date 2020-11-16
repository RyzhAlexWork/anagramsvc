// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/api/v1"
)

//easyjson:json
type getAnagramsResponse struct {
	v1.GetAnagramsResponse
}

// GetAnagramsTransport transport interface
type GetAnagramsTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, searchWord *string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetAnagramsResponse, err error)
}

//easyjson:skip
type getAnagramsTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *getAnagramsTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, searchWord *string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	if searchWord != nil {
		r.URI().QueryArgs().Set("word", *searchWord)
	}

	return
}

// DecodeResponse method for encoding response on server side
func (t *getAnagramsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.GetAnagramsResponse, err error) {
	if r.StatusCode() != v1.HTTPStatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse getAnagramsResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	response = theResponse.GetAnagramsResponse

	return
}

// NewGetAnagramsTransport the transport creator for http requests
func NewGetAnagramsTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) GetAnagramsTransport {
	return &getAnagramsTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

//easyjson:skip
type loadWordsRequest struct {
	LoadWords *[]string `json:"loadWords"`
}

//easyjson:json
type loadWordsResponse struct {
	v1.LoadWordsResponse
}

// LoadWordsTransport transport interface
type LoadWordsTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, loadWords *[]string) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.LoadWordsResponse, err error)
}

//easyjson:skip
type loadWordsTransport struct {
	errorProcessor errorProcessor
	pathTemplate   string
	method         string
}

// EncodeRequest method for decoding requests on server side
func (t *loadWordsTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, loadWords *[]string) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	return
}

// DecodeResponse method for encoding response on server side
func (t *loadWordsTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.LoadWordsResponse, err error) {
	if r.StatusCode() != v1.HTTPStatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}

	var theResponse loadWordsResponse
	if err = theResponse.UnmarshalJSON(r.Body()); err != nil {
		return
	}

	response = theResponse.LoadWordsResponse

	return
}

// NewLoadWordsTransport the transport creator for http requests
func NewLoadWordsTransport(
	errorProcessor errorProcessor,
	pathTemplate string,
	method string,
) LoadWordsTransport {
	return &loadWordsTransport{
		errorProcessor: errorProcessor,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}
