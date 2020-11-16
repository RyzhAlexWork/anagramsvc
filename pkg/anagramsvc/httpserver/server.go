// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"

	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/api/v1"
)

type service interface {
	GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error)
	LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error)
}

type getAnagrams struct {
	transport      GetAnagramsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getAnagrams) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		searchWord *string
		response   v1.GetAnagramsResponse
		err        error
	)
	searchWord, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err = s.service.GetAnagrams(ctx, searchWord)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetAnagrams the server creator
func NewGetAnagrams(transport GetAnagramsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getAnagrams{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type loadWords struct {
	transport      LoadWordsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *loadWords) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		loadWords *[]string
		response  v1.LoadWordsResponse
		err       error
	)
	loadWords, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err = s.service.LoadWords(ctx, loadWords)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewLoadWords the server creator
func NewLoadWords(transport LoadWordsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := loadWords{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
