package httphandlers

import (
	"app/modules/stats/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetMostAppearError(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	limit := c.Param("limit")
	result, err := repositories.GetMostAppearError(page, 10, "api/v1/stats/"+ord+"/"+limit, ord, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMostCreatedTagByCategory(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	limit := c.Param("limit")
	result, err := repositories.GetMostCreatedTagByCategory(page, 10, "api/v1/stats/tagcat/"+ord+"/"+limit, ord, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
