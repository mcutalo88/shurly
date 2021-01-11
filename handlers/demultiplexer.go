package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ShurlyLink struct {
	Vanity string   `json:"vanity,omitempty"`
	Links  []string `json:"links"`
}

// CreateShurlyUrl
func CreateShurlyUrl(w http.ResponseWriter, r *http.Request) {
	var u ShurlyLink

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Created: [%v]", u)
}
