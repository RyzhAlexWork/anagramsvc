// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodGetAnagrams = "GET"
	uriPathGetAnagrams    = "/get"
	httpMethodLoadWords   = "POST"
	uriPathLoadWords      = "/load"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, decodeTypeIntErrorCreator errorCreator, errorProcessor errorProcessor) {

	getAnagramsTransport := NewGetAnagramsTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGetAnagrams, uriPathGetAnagrams, NewGetAnagrams(getAnagramsTransport, svc, errorProcessor))

	loadWordsTransport := NewLoadWordsTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodLoadWords, uriPathLoadWords, NewLoadWords(loadWordsTransport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
