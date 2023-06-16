package main

import (
	"app/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
