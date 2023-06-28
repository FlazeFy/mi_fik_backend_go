package routes

import (
	middleware "app/middlewares/jwt"
	authhandlers "app/modules/auth/http_handlers"
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

	// =============== Public routes ===============
	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)

	// Tag
	e.GET("api/v1/tag", taghandlers.GetAllActiveTag)
	e.GET("api/v1/tag/:category", taghandlers.GetAllActiveTagByCategory)
	e.GET("api/v1/trash/tag", taghandlers.GetAllTrashTag, middleware.CustomJWTAuth)

	// Dictionary
	e.GET("api/v1/dct", dcthandlers.GetAllActiveDictionaries)
	e.GET("api/v1/dct/color", dcthandlers.GetAllRecentColor)
	e.GET("api/v1/trash/dct", dcthandlers.GetAllTrashDictionaries, middleware.CustomJWTAuth)

	// =============== Private routes ===============

	return e
}
