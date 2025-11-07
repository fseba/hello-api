// Package rest handles HTTP requests for translations.
package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

const defaultLanguage = "english"

type Translator interface {
	Translate(word, language string) string
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler will translate calls for caller.
type TranslateHandler struct {
	service Translator
}

// NewTranslateHandler will create a new instance of the handler using a translator service.
func NewTranslateHandler(service Translator) *TranslateHandler {
	return &TranslateHandler{service: service}
}

// TranslateHandler will take a given request with a path value of the
// word to be translated and a query parameter of language to translate to.
func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := t.service.Translate(word, language)
	if translation == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
