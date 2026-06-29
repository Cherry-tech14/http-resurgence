package main

import (
	"fmt"
	"net/http"
)

// func methodInspectorHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {

// 	} else if r.Method != "POST" {
// 		return
// 	}
// 	fmt.Fprintf(w, "You made a %s request.", r.Method)

// }

// Second way

func methodInspectorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You made a %v request.", r.Method)
}
