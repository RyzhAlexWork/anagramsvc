// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"anagramsvc/pkg/api/v1"
)

// Options ...
var (
	GetAnagrams = option{}
	LoadWords   = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service ...
type Service interface {
	GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error)
	LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error)
}

type client struct {
	cli                  *fasthttp.HostClient
	transportGetAnagrams GetAnagramsTransport
	transportLoadWords   LoadWordsTransport
	options              map[interface{}]Option
}

// GetAnagrams ...
func (s *client) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetAnagrams]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetAnagrams.EncodeRequest(ctx, req, searchWord); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetAnagrams.DecodeResponse(ctx, res)
}

// LoadWords ...
func (s *client) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[LoadWords]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportLoadWords.EncodeRequest(ctx, req, loadWords); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportLoadWords.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportGetAnagrams GetAnagramsTransport,
	transportLoadWords LoadWordsTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli:                  cli,
		transportGetAnagrams: transportGetAnagrams,
		transportLoadWords:   transportLoadWords,
		options:              options,
	}
}
