// Package anagramsvc ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package anagramsvc

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"anagramsvc/pkg/api/v1"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

// GetAnagrams ...
func (s *loggingMiddleware) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetAnagrams",
			"searchWord", searchWord,

			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetAnagrams(ctx, searchWord)
}

// LoadWords ...
func (s *loggingMiddleware) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "LoadWords",
			"loadWords", loadWords,

			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.LoadWords(ctx, loadWords)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
