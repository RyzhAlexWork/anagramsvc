// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/api/v1"
)

var (
	emptyBytes = []byte("")
)

//easyjson:json
type getAnagramsResponse struct {
	v1.GetAnagramsResponse
}

// GetAnagramsTransport transport interface
type GetAnagramsTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (searchWord *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetAnagramsResponse) (err error)
}

type getAnagramsTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getAnagramsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (searchWord *string, err error) {

	searchWord = ptr(ctx.QueryArgs().Peek("word"))

	return
}

// EncodeResponse method for encoding response on server side
func (t *getAnagramsTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetAnagramsResponse) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse getAnagramsResponse

	theResponse.GetAnagramsResponse = response

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewGetAnagramsTransport the transport creator for http requests
func NewGetAnagramsTransport(

	encodeJSONErrorCreator errorCreator,

) GetAnagramsTransport {
	return &getAnagramsTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

//easyjson:json
type loadWordsResponse struct {
	v1.LoadWordsResponse
}

// LoadWordsTransport transport interface
type LoadWordsTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (loadWords *[]string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.LoadWordsResponse) (err error)
}

type loadWordsTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *loadWordsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (loadWords *[]string, err error) {
	err = json.Unmarshal(r.Body(), &loadWords)
	return
}

// EncodeResponse method for encoding response on server side
func (t *loadWordsTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.LoadWordsResponse) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse loadWordsResponse

	theResponse.LoadWordsResponse = response

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewLoadWordsTransport the transport creator for http requests
func NewLoadWordsTransport(

	encodeJSONErrorCreator errorCreator,

) LoadWordsTransport {
	return &loadWordsTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
