package v1

import (
	"net/http"
)

// Const ...
const (
	HTTPStatusOK         = http.StatusOK
	HTTPStatusBadRequest = http.StatusBadRequest

	RegularExpression = `^[a-zA-Zа-яА-Я]+$`
	LoadMessage       = "Words uploaded successfully"
)

// AdditionalErrors ...
type AdditionalErrors struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

// GetAnagramsResponse ...
type GetAnagramsResponse struct {
	Data             *[]string         `json:"data"`
	Error            bool              `json:"error"`
	ErrorText        string            `json:"errorText"`
	AdditionalErrors *AdditionalErrors `json:"additionalErrors"`
}

// LoadWordsResponse ...
type LoadWordsResponse struct {
	Data             string            `json:"data"`
	Error            bool              `json:"error"`
	ErrorText        string            `json:"errorText"`
	AdditionalErrors *AdditionalErrors `json:"additionalErrors"`
}
