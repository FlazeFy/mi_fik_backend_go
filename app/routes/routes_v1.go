package routes

import (
	dcthandlers "app/modules/systems/http_handlers"
	taghandlers "app/modules/tags/http_handlers"
	"net/http"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Mi-FIK")
	})

	// Tag
	e.GET("api/v1/tag", taghandlers.GetAllActiveTag)
	e.GET("api/v1/tag/:category", taghandlers.GetAllActiveTagByCategory)
	e.GET("api/v1/trash/tag", taghandlers.GetAllTrashTag)

	// Dictionary
	e.GET("api/v1/dct", dcthandlers.GetAllActiveDictionaries)
	e.GET("api/v1/trash/dct", dcthandlers.GetAllTrashDictionaries)

	return e
}
