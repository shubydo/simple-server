package main

import (
	"fmt"
	"net/http"
)

func initHandlers() {
	http.HandleFunc("/hello", handleHello)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello!")
		return
	}

	// Return "Method not allowed" if not a GET request
	fmt.Fprintf(w, "Method not allowed")
}
