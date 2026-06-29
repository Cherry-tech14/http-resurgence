package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "stranger"
	}
	fmt.Fprintf(w, "Greetings, %s!", name)
}
func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/v1/ping", pingHandler)

	apiMux.HandleFunc("/v1/greet", greetHandler)
	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("server running on port :8080")

	http.ListenAndServe(":8080", mainMux)
}
