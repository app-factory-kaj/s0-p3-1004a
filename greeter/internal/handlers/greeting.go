package handlers

import (
	"encoding/json"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

type Error struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
	MoreInfo    string `json:"moreInfo,omitempty"`
}

// GetGreeting handles GET /hello, returning a JSON greeting personalized by
// the optional "name" query parameter and defaulting to "World" when it is
// omitted or empty.
func GetGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Greeting{Message: "Hello, " + name + "!"})
}
