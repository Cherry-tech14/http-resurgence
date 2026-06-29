package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("server running on port :8080")
	http.HandleFunc("/method-inspector", methodInspectorHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe(":8080", nil)
}
