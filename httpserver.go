package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GreetingRequest struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello %s", name)
		return
	}

	if r.Method == http.MethodPost {
		var req GreetingRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if req.name == "" {
			req.name = "Guest"
		}
		fmt.Fprintf(w, "Hello %s", req.name)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
