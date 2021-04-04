package router

import (
	"oauth2-authorization/models"
	"oauth2-authorization/utility"

	"github.com/labstack/echo/v4"
)

const (
	prefix = "/api/v1"
)

func InitRouter() (e *echo.Echo) {
	router := echo.New()
	AddUserRouter(prefix+"/user", router)
	AddOauth2Router(prefix+"/oauth2", router)

	utility.Log(models.LogLevelDebug, nil, nil, "Router Init Complete!!")

	return router
}
