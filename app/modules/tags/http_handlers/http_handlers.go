package httphandlers

import (
	"app/modules/tags/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllActiveTag(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllTag(page, 10, "api/v1/tag", "active")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllActiveTagByCategory(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	category := c.Param("category")
	result, err := repositories.GetAllTagByCategory(page, 10, "api/v1/tag/:"+category, "active", category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTrashTag(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllTag(page, 10, "api/v1/tag", "trash")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
