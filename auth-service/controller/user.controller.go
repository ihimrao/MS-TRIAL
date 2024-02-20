package controller

import (
	database "auth-service/db"
	middlewares "auth-service/handlers"
	user_model "auth-service/models"
	"auth-service/utils"
	"context"
	"encoding/json"
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

var AuthorizeUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var uStr = r.Header["uid"][0]
	userId, err := primitive.ObjectIDFromHex(uStr)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Un-Authorized"))
	}
	var user user_model.User
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Un-Authorized"))
	}
	w.WriteHeader(http.StatusCreated)
	res := bson.M{"id": userId, "Authorized": true, "Permissions": []string{""}}
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusOK, "Authorized", res))
})
