package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	fmt.Fprintln(c.Response().Writer, "Hello World!!")

	return nil
}
