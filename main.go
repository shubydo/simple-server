package main

import (
	"fmt"
	"log"
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

func main() {
	initHandlers()

	port := 8080
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("Error occured: %v\n", err)
	}
}
