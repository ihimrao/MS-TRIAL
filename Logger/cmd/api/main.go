package main

import (
	"context"
	"fmt"
	"log"
	"mstrail/data"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Models data.Models
}

const (
	port     = "8080"
	mongoURL = "mongodb+srv://ihimrao:sciJz7anvmj2vDWB@lms.e5ahgmo.mongodb.net/?retryWrites=true&w=majority"
)

func main() {
	mongoClient, err := ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to Mongo", err)
	}
	fmt.Println("hello")
	fmt.Println("hello")
	app := &Config{
		Models: data.New(mongoClient),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.Routes(),
	}

	log.Printf("Starting Logger Service on PORT %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error listening and serving", err)
	}
}

func ConnectToDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting database", err)
		return nil, err
	}
	fmt.Println("Connected to database")
	return c, nil
}
