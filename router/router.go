package router

import (
	"oauth2-authorization/utility"

	"github.com/gorilla/mux"
)

const (
	prefix = "/api/v1"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router = AddUserRouter(prefix+"/user", router)
	router = AddOauth2Router(prefix+"/oauth2", router)

	utility.Log("Debug", nil, nil, "Router Init Complete!!")

	return router
}
