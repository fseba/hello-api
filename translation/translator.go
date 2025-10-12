package translation

import "strings"

func Translate(word, language string) string {
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
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.TrimSpace(w)
	return strings.ToLower(w)
}
