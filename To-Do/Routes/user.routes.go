package route

import (
	controller "go-base-fs/controllers"
	middlewares "go-base-fs/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/createUser", controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/ping", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("pong!"))
	}).Methods("GET")

	// todo's protected route
	router.HandleFunc("/todo", middlewares.IsAuthorized(controller.AddToDo)).Methods("POST")
	router.HandleFunc("/todo/{id}", middlewares.IsAuthorized(controller.UpdateToDo)).Methods("PUT")
	router.HandleFunc("/todo", middlewares.IsAuthorized(controller.GetAllTodo)).Methods("GET")
	router.HandleFunc("/todo/{id}", middlewares.IsAuthorized(controller.DeleteToDo)).Methods("DELETE")
	router.HandleFunc("/todo/{id}", middlewares.IsAuthorized(controller.GetTodo)).Methods("GET")
	return router
}
