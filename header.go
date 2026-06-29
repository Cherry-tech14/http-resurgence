package main

import (
	"fmt"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		contentType = "Content-Type not provided"

	}
	fmt.Fprintf(w, "Token received: %s\nContent-Type: %s", token, contentType)
}
