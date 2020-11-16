package v1

// Err ...
const (
	MissingWordListError = "need to load word list"
	OnlyWordError        = "only words can be loaded"
	SearchWordError      = "query parameter must be a word"

	MissingSearchWordParamError = "2: wrong query parameter: searchWord"

	MissingLoadWordsParamError = "3: missing required parameter: loadWords"
)

// ErrMap ...
var (
	ErrMap = map[string]string{
		"1": "service error",
		"2": "wrong query parameter",
		"3": "missing required parameter",
	}
)
