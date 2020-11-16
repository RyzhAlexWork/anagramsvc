package anagramsvc

import (
	"context"
	"fmt"
	"regexp"

	"anagramsvc/pkg/api/v1"
)

// validatorMiddleware wraps Service and validate request
type validatorMiddleware struct {
	next         Service
	debug        bool
	errorCreator errorCreator
}

// GetAnagrams ...
func (m *validatorMiddleware) GetAnagrams(ctx context.Context, searchWord *string) (response v1.GetAnagramsResponse, err error) {
	if searchWord == nil {
		err = m.errorCreator(v1.HTTPStatusBadRequest, v1.MissingSearchWordParamError)
		return
	}
	isWord, _ := regexp.MatchString(v1.RegularExpression, *searchWord)
	if !isWord {
		err = fmt.Errorf(v1.SearchWordError)
		return
	}
	return m.next.GetAnagrams(ctx, searchWord)
}

// LoadWords ...
func (m *validatorMiddleware) LoadWords(ctx context.Context, loadWords *[]string) (response v1.LoadWordsResponse, err error) {
	if loadWords == nil {
		err = m.errorCreator(v1.HTTPStatusBadRequest, v1.MissingLoadWordsParamError)
		return
	}
	for _, word := range *loadWords {
		isWord, _ := regexp.MatchString(v1.RegularExpression, word)
		if !isWord {
			err = fmt.Errorf(v1.OnlyWordError)
			return
		}
	}
	return m.next.LoadWords(ctx, loadWords)
}

// NewValidatorMiddleware ...
func NewValidatorMiddleware(
	next Service,
	debug bool,
	errorCreator errorCreator,
) Service {
	return &validatorMiddleware{
		next:         next,
		debug:        debug,
		errorCreator: errorCreator,
	}
}
