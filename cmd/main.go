package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fseba/hello-api/config"
	"github.com/fseba/hello-api/handlers"
	"github.com/fseba/hello-api/handlers/rest"
	"github.com/fseba/hello-api/translation"
)

func main() {
	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}
	translateHandler := rest.NewTranslateHandler(translationService, cfg.DefaultLanguage)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

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
