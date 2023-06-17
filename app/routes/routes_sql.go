package routes

import (
	httphandlers "app/modules/tags/http_handlers"
	"net/http"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Mi-FIK")
	})

	//Tag
	e.GET("api/v1/tag", httphandlers.GetAllActiveTag)
	e.GET("api/v1/trash/tag", httphandlers.GetAllTrashTag)

	return e
}
