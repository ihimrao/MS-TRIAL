package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Context-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_ = app.WriteJson(w, http.StatusOK, "logger-service")
	})
	router.Post("/log", app.writeLog)
	return router
}
