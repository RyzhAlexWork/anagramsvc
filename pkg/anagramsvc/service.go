package anagramsvc

import (
	"context"

	"anagramsvc/pkg/api/rest"
	"anagramsvc/pkg/api/v1"
)

type errorCreator func(status int, format string, v ...interface{}) error

type processor interface {
	GetAnagrams(ctx context.Context, req rest.Request) (response v1.GetAnagramsResponse, err error)
	LoadWords(ctx context.Context, req rest.Request) (response v1.LoadWordsResponse, err error)
}

// Service ...
// @gtg http-server log metrics mock http-client http-errors
type Service interface {
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /get
	// @gtg http-server-query word={searchWord}
	// @gtg http-server-response-body response
	// @gtg http-server-response-status v1.HTTPStatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg log-ignore response
	GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error)
	// @gtg http-server-method POST
	// @gtg http-server-uri-path /load
	// @gtg http-server-json-tag loadWords loadWords
	// @gtg http-server-response-body response
	// @gtg http-server-response-status v1.HTTPStatusOK
	// @gtg http-server-response-content-type application/json
	// @gtg log-ignore response
	LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error)
}

type service struct {
	processor processor
}

// GetAnagrams ...
func (s *service) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	req := rest.NewRequest().SetSearchWord(searchWord)

	return s.processor.GetAnagrams(ctx, req)
}

// LoadWords ...
func (s *service) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	req := rest.NewRequest().SetLoadWords(loadWords)

	return s.processor.LoadWords(ctx, req)
}

// NewService ...
func NewService(processor processor) Service {
	return &service{
		processor: processor,
	}
}
