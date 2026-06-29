package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	codeString := r.URL.Query().Get("code")

	if codeString == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	code, err := strconv.Atoi(codeString)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}
	if code < 100 || code > 599 {
		http.Error(w, "code must be a valid HTTP status code (100-599)", http.StatusBadRequest)

	}
	w.WriteHeader(code)
	fmt.Fprintf(w, "Responding with status %d", code)
}
