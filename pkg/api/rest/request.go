package rest

type request struct {
	searchWord *string
	loadWords  *[]string
}

// Request ...
type Request interface {
	SearchWord() *string
	LoadWords() *[]string

	SetSearchWord(searchWord *string) Request
	SetLoadWords(loadWords *[]string) Request
}

// SearchWord ...
func (r *request) SearchWord() *string {
	return r.searchWord
}

// LoadWords ...
func (r *request) LoadWords() *[]string {
	return r.loadWords
}

// SetSearchWord ...
func (r *request) SetSearchWord(searchWord *string) Request {
	r.searchWord = searchWord
	return r
}

// SetLoadWords ...
func (r *request) SetLoadWords(loadWords *[]string) Request {
	r.loadWords = loadWords
	return r
}

// NewRequest ...
func NewRequest() Request {
	return &request{}
}
