package main

import (
	route "auth-service/Routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}
	color.Cyan("🌏 Server running on localhost:" + os.Getenv("PORT"))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := route.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
