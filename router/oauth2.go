package router

import (
	"oauth2-authorization/controllers"

	"github.com/labstack/echo/v4"
)

func AddOauth2Router(prefix string, e *echo.Echo) {
	e.GET(prefix+"/", controllers.HelloWorld)
}
