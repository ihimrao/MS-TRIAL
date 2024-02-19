package controller

import (
	"context"
	"encoding/json"
	database "go-base-fs/db"
	middlewares "go-base-fs/handlers"
	user_model "go-base-fs/models"
	"go-base-fs/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client = database.DatabaseConnection()
var userCollection = client.Database(utils.GetEnvVar("DB_NAME")).Collection("USER")

var Login = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var user user_model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	filter := bson.M{"email": user.Email}
	var u bson.M
	if err := userCollection.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			response.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	valid := utils.CheckPasswordHash(user.Password, u["password"].(string))
	if valid {
		validToken, err := middlewares.GenerateJWT(u["_id"].(primitive.ObjectID).Hex())
		if err != nil {
			json.NewEncoder(response).Encode(utils.ErrorResponse(http.StatusBadRequest, ""))
		}
		json.NewEncoder(response).Encode(utils.SuccessResponse(http.StatusOK, "Successfully Logged In", validToken))
	}
})

var CreateUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user user_model.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.Password = utils.HashPassword(user.Password)
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal("Error creating User")
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusOK, "User Created Successfully", result))
})
