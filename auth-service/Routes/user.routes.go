package route

import (
	"auth-service/controller"
	middlewares "auth-service/handlers"
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
	router.HandleFunc("/authorize", middlewares.IsAuthorized(controller.AuthorizeUserHandler)).Methods("POST")
	return router
}
