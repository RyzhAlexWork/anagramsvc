// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"

	"github.com/stretchr/testify/mock"

	"anagramsvc/pkg/api/v1"
)

// MockService ...
type MockService struct {
	mock.Mock
}

// GetAnagrams ...
func (s *MockService) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	args := s.Called(context.Background(), searchWord)
	return args.Get(0).(v1.GetAnagramsResponse), args.Error(1)
}

// LoadWords ...
func (s *MockService) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	args := s.Called(context.Background(), loadWords)
	return args.Get(0).(v1.LoadWordsResponse), args.Error(1)
}
