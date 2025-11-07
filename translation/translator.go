// Package translation provides a function to translate the word "hello" into different languages.
package translation

import "strings"

// StaticService has data that does not change.
type StaticService struct{}

// NewStaticService creates a new instance of StaticService that uses static data.
func NewStaticService() *StaticService {
	return &StaticService{}
}

// Translate a given word to the passed language.
func (s *StaticService) Translate(word, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)
	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	case "finnish":
		return "hei"
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.TrimSpace(w)
	return strings.ToLower(w)
}
