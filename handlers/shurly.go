package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mcutalo88/shurly/pkg/types"
)

type Shurly struct {
	Vanity string   `json:"Vanity,omitempty"`
	Links  []string `json:"links"`
}

func CreateShurlyUrl(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(types.DatabaseContext).(*sql.DB)

	var shurly Shurly

	err := json.NewDecoder(r.Body).Decode(&shurly)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Created: [%v]", shurly)

	insert := `
		INSERT INTO shurly_links (id, url, links) VALUES ($1, $2, $3)
	`

	shurlyLinks := strings.Join(shurly.Links[:], ",")

	_, err = db.ExecContext(r.Context(),
		insert,
		shurly.Vanity,
		fmt.Sprintf("http://localhost:8000/%s", shurly.Vanity),
		fmt.Sprintf("{%s}", shurlyLinks))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
