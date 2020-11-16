// Package anagramsvc ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package anagramsvc

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"

	"anagramsvc/pkg/api/v1"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

// GetAnagrams ...
func (s *instrumentingMiddleware) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	defer s.recordMetrics("GetAnagrams", time.Now(), err)
	return s.svc.GetAnagrams(ctx, searchWord)
}

// LoadWords ...
func (s *instrumentingMiddleware) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	defer s.recordMetrics("LoadWords", time.Now(), err)
	return s.svc.LoadWords(ctx, loadWords)
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
