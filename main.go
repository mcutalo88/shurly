package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcutalo88/shurly/handlers"
	"github.com/mcutalo88/shurly/pkg/config"
	"github.com/mcutalo88/shurly/pkg/db"
	"github.com/mcutalo88/shurly/pkg/types"
	"go.uber.org/zap"
)

func main() {
	cfg := config.ReadConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("Fatal not setup logger: %v", err))
	}

	db := db.New(cfg)
	defer db.Close()

	router := mux.NewRouter()

	// TODO: Refactor logging injection later.
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), types.DatabaseContext, db)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/shurly", handlers.CreateShurlyUrl).Methods("POST")

	logger.Info("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
