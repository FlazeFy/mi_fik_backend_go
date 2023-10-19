package routes

import (
	middlewares "app/middlewares/jwt"
	authhandlers "app/modules/auth/http_handlers"
	statshandlers "app/modules/stats/http_handlers"
	dcthandlers "app/modules/systems/http_handlers"
	taghandlers "app/modules/tags/http_handlers"
	userhandlers "app/modules/users/http_handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NON ORM
func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

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

	// Dictionary
	e.GET("api/v1/dct", dcthandlers.GetAllActiveDictionaries)
	e.GET("api/v1/dct/color", dcthandlers.GetAllRecentColor)

	// Stats
	e.GET("api/v1/stats/:ord/:limit", statshandlers.GetMostAppearError)

	// =============== Private routes ===============
	// Tag
	e.GET("api/v1/trash/tag", taghandlers.GetAllTrashTag, middlewares.CustomJWTAuth)

	// Dictionary
	e.GET("api/v1/trash/dct", dcthandlers.GetAllTrashDictionaries, middlewares.CustomJWTAuth)

	// User
	e.GET("api/v1/user", userhandlers.GetMyProfile, middlewares.CustomJWTAuth)

	return e
}
