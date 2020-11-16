// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	httpMethodGetAnagrams    = "GET"
	uriPathClientGetAnagrams = "/get"
	httpMethodLoadWords      = "POST"
	uriPathClientLoadWords   = "/load"
)

type errorProcessor interface {
	Decode(r *fasthttp.Response) error
}

// New ...
func New(
	serverURL string,
	maxConns int,
	errorProcessor errorProcessor,
	options map[interface{}]Option,
) (client Service, err error) {
	parsedServerURL, err := url.Parse(serverURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse server url")
		return
	}
	transportGetAnagrams := NewGetAnagramsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientGetAnagrams,
		httpMethodGetAnagrams,
	)
	transportLoadWords := NewLoadWordsTransport(
		errorProcessor,
		parsedServerURL.Scheme+"://"+parsedServerURL.Host+uriPathClientLoadWords,
		httpMethodLoadWords,
	)

	client = NewClient(
		&fasthttp.HostClient{
			Addr:     parsedServerURL.Host,
			MaxConns: maxConns,
		},
		transportGetAnagrams,
		transportLoadWords,
		options,
	)
	return
}
