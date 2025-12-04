// Package faas implements the function-as-a-service entry points.
package faas

import (
	"net/http"

	"github.com/fseba/hello-api/handlers/rest"
	"github.com/fseba/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService, "english")
	translateHandler.TranslateHandler(w, r)
}
