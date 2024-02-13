package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/ping"))
	return router
}
