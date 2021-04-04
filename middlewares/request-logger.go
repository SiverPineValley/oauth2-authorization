package middlewares

import (
	"oauth2-authorization/models"
	"oauth2-authorization/utility"

	"github.com/labstack/echo/v4"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utility.Log(models.LogLevelInfo, c.Response().Writer, c.Request())
		return next(c)
	}
}
