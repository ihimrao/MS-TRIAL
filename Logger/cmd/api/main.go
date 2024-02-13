package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const (
	port = "8080"
)

func main() {
	app := &Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.Routes(),
	}

	log.Println("Starting logger service")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error listening and serving", err)
	}
}
