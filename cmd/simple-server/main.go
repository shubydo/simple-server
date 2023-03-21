package main

import (
	"github.com/shubydo/simple-server/internal/server"
)

func main() {
	s := server.NewServer()
	s.ServeHTTP(nil, nil)
}
