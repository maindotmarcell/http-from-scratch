package main

import (
	"log"

	"github.com/maindotmarcell/http-from-scratch/internal/handler"
	"github.com/maindotmarcell/http-from-scratch/internal/server"
)

func main() {

	s := server.NewHTTPServer("0.0.0.0:3000")

	// Assign handlers to paths here
	s.Router.HandleGet("/", handler.HandleRoot)

	// Starting the server
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
