package middlewares

import (
	"app/modules/auth/models"
	"app/modules/auth/repositories"
	"app/packages/helpers/auth"
	"app/packages/helpers/generator"
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

func CheckLogin(c echo.Context, body models.UserLogin) (response.Response, error) {
	var res response.Response
	timeExpStr, err := strconv.Atoi(auth.GetJWTConfiguration("exp"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		return res, err
	}

	duration := time.Duration(timeExpStr) * time.Second

	authResult, err, ctx := repositories.PostUserAuth(body.Username, body.Password)
	// Response
	if err == nil && !authResult {
		res.Status = http.StatusUnprocessableEntity
		res.Message = ctx
		return res, err
	}

	if err != nil {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return res, err
	}

	if !authResult {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return res, err
	}

	// Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = body.Username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * duration).Unix()

	// Response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
	}

	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("login", "", true)
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}
