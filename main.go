package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mcutalo88/shurly/handlers"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	router := mux.NewRouter()

	// TODO: Refactor logging injection later.
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "log", logger)
			h.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/shurly", handlers.CreateShurlyUrl).Methods("POST")

	logger.Info("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
