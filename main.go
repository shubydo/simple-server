package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initHandlers()

	port := 8080
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("Error occured: %v\n", err)
	}
}
