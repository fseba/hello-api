// Package faas implements the function-as-a-service entry points.
package faas

import (
	"net/http"

	"github.com/fseba/hello-api/handlers/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
