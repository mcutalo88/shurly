package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mcutalo88/shurly/handlers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	viper.SetConfigFile("shurly-config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

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
