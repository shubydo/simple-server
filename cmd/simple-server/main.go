package main

import (
	"log"
	"net/http"

	"github.com/shubydo/simple-server/internal/server"
)

func main() {
	s := server.NewServer()
	log.Fatalln(http.ListenAndServe(":8080", s))
}
