package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	if language == "" {
		http.Error(w, "language is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s, you are coding in %v!", username, language)
}
