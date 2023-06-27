package middlewares

import (
	"app/modules/auth/models"
	"app/modules/auth/repositories"
	"app/packages/helpers/auth"
	"app/packages/helpers/response"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func CheckLogin(c echo.Context, body models.UserLogin) error {
	var res response.Response
	timeExpStr, _ := strconv.Atoi(auth.GetJWTConfiguration("exp"))
	duration := time.Duration(timeExpStr) * time.Second

	authResult, err, ctx := repositories.PostUserAuth(body.Username, body.Password)
	if err != nil {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return c.JSON(http.StatusUnauthorized, res)
	}

	if !authResult {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return c.JSON(http.StatusUnauthorized, res)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = body.Username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * duration).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Status = http.StatusOK
	res.Data = map[string]string{
		"token": t,
	}

	return c.JSON(http.StatusOK, res)
}
