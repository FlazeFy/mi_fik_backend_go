package httphandlers

import (
	"app/modules/users/repositories"
	"app/packages/helpers/auth"
	"net/http"

	"github.com/labstack/echo"
)

func GetMyProfile(c echo.Context) error {
	status, token := auth.GetTokenHeader(c)
	if !status {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": token})
	}

	result, err := repositories.GetMyProfile(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
