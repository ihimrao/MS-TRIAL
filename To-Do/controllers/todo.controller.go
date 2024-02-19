package controller

import (
	"context"
	"encoding/json"
	model "go-base-fs/models"
	"go-base-fs/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todoCollection = client.Database(utils.GetEnvVar("DB_NAME")).Collection("TODO")

var AddToDo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var uStr = r.Header["uid"][0]
	userId, err := primitive.ObjectIDFromHex(uStr)
	if err != nil {
		log.Fatal("Malformed id", err)
	}
	var addToDo model.Todo
	addToDo.UserID = userId
	json.NewDecoder(r.Body).Decode(&addToDo)
	res, err := todoCollection.InsertOne(context.Background(), addToDo)
	if err != nil {
		log.Fatal("Error inserting todo", err)
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Error in payload"))
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusCreated, "Todo Created Successfully", res))
})

var GetTodo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	user := r.Header["uid"][0]
	userId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		log.Fatal("Error fetching records", err)
	}
	todoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Error fetching records", err)
	}
	var todo bson.M
	todoCollection.FindOne(context.TODO(), bson.M{"_id": todoId, "userId": userId}).Decode(&todo)
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusOK, "Successfully fetched", todo))
})

var GetAllTodo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var users []bson.M
	var uStr = r.Header["uid"][0]
	userId, err := primitive.ObjectIDFromHex(uStr)
	if err != nil {
		log.Fatal("Malformed id", err)
	}
	cursor, err := todoCollection.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		log.Fatal("Error inserting todo", err)
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Error Fetching Records"))

	}
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("Error inserting todo", err)
			json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Unexpected Error"))

		}
		users = append(users, user)
	}
	defer cursor.Close(context.Background())
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusCreated, "Users fetched Successfully", users))
})

var DeleteToDo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Please pass a valid todo id"))
	}
	var u bson.M
	userId := r.Header["uid"][0]
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	todoCollection.FindOne(context.TODO(), bson.M{"_id": todoID}).Decode(&u)
	_, err = todoCollection.DeleteOne(context.Background(), bson.M{"_id": todoID, "userId": userIdObj})
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Unable to delete todo"))
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusOK, "Deleted Successfully", bson.M{}))

})

var UpdateToDo = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var updatedTodo model.Todo
	params := mux.Vars(r)
	todoID := params["id"]
	userId := r.Header["uid"][0]
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Please pass a valid todo id"))
	}
	todoObjId, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Please pass a valid todo id"))
	}
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		log.Fatal("Please provide a valid data")
	}
	update := bson.M{
		"$set": bson.M{
			"completed":   updatedTodo.Completed,
		},
	}
	nUpdated, err := todoCollection.UpdateOne(context.Background(), bson.M{"_id": todoObjId, "userId": userObjId}, update)
	if nUpdated.ModifiedCount > 0 {
		json.NewEncoder(w).Encode(utils.SuccessResponse(http.StatusOK, "Updated Successfully", bson.M{}))
	} else {
		json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusBadRequest, "Cannot find existing data"))
	}
})
