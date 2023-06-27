package httphandlers

import (
	middlewares "app/middlewares/jwt"
	"app/modules/auth/models"
	"app/modules/auth/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func PostLoginUser(c echo.Context) error {
	var body models.UserLogin
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}

	result := middlewares.CheckLogin(c, body)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func PostRegister(c echo.Context) error {
	var body models.UserRegister
	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}

	result, err := repositories.PostUserRegister(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
