package main

import (
	"log"

	"github.com/maindotmarcell/http-from-scratch/internal/handler"
	"github.com/maindotmarcell/http-from-scratch/internal/server"
)

func main() {

	s := server.NewHTTPServer("0.0.0.0:3000")

	s.Router.HandleGet("/", handler.HandleRoot)

	log.Fatal(s.Start())

}
