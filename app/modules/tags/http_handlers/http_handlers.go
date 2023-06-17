package httphandlers

import (
	"app/modules/tags/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllActiveTag(c echo.Context) error {
	result, err := repositories.GetAllTag(1, 10, "api/v1/tag", "active")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTrashTag(c echo.Context) error {
	result, err := repositories.GetAllTag(1, 10, "api/v1/tag", "trash")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
