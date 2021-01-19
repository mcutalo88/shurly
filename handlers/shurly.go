package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mcutalo88/shurly/internal/types"
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

	_, err = db.ExecContext(r.Context(),
		"INSERT INTO shurly_links (id, url, links) VALUES ($1, $2, $3)",
		shurly.Vanity,
		fmt.Sprintf("http://localhost:8000/%s", shurly.Vanity),
		fmt.Sprintf("{%s}", strings.Join(shurly.Links[:], ",")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
