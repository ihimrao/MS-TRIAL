package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Models
}

const (
	port = "8080"
)

var mongoURL string

func main() {
	// mongoClient, err := ConnectToDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL = os.Getenv("MONGO_URI")
	if err != nil {
		log.Fatal("Error connecting to Mongo", mongoURL)
	}
	// app := &Config{
	// 	Models: data.New(mongoClient),
	// }

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}),
	}

	// log.Printf("Starting Logger Service on PORT %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error listening and serving", err)
	}
}

// func ConnectToDB() (*mongo.Client, error) {
// 	clientOptions := options.Client().ApplyURI(mongoURL)
// 	c, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal("Error connecting database", err)
// 		return nil, err
// 	}
// 	fmt.Println("Connected to database")
// 	return c, nil
// }
