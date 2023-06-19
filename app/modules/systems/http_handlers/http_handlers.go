package httphandlers

import (
	"app/modules/systems/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllActiveDictionaries(c echo.Context) error {
	result, err := repositories.GetAllDictionary(1, 10, "api/v1/dct", "active")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTrashDictionaries(c echo.Context) error {
	result, err := repositories.GetAllDictionary(1, 10, "api/v1/dct", "trash")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
