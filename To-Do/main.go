package main

import (
	route "go-base-fs/Routes"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	color.Cyan("🌏 Server running on localhost:" + os.Getenv("PORT"))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := route.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
