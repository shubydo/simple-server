package main

import (
	"log"
)

func main() {
	s := server.NewServer(
		server.WithPort(8080),
	)

	log.Fatal(s.ListenAndServe())
}
