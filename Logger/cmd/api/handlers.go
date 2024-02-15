package main

import (
	"encoding/json"
	"mstrail/data"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) writeLog(w http.ResponseWriter, r *http.Request) {
	// var logPayload JSONPayload

	// err := app.ReadJson(w, r, logPayload)
	// if err != nil {
	// 	app.ErrorJson(w, err, http.StatusBadRequest)
	// }
	var Data data.LogEntry
	json.NewDecoder(r.Body).Decode(&Data)
	err := app.Models.LogEntry.Insert(Data)
	if err != nil {
		app.ErrorJson(w, err, http.StatusBadRequest)
	}
	app.WriteJson(w, http.StatusCreated, bson.M{
		"Error":   false,
		"Message": "Logged",
	})

}
