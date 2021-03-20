package router

import (
	"oauth2-authorization/controllers"

	"github.com/gorilla/mux"
)

func AddUserRouter(prefix string, router *mux.Router) *mux.Router {
	router.HandleFunc(prefix+"/", controllers.HelloWorld).Methods("GET")

	return router
}
