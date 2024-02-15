package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) WriteJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

func (app *Config) ErrorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.WriteJson(w, statusCode, payload)
}
