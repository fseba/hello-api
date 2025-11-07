package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fseba/hello-api/handlers"
	"github.com/fseba/hello-api/handlers/rest"
	"github.com/fseba/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("listening on %s\n", addr)

	log.Fatal(server.ListenAndServe())
}
