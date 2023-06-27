package auth

import (
	"app/configs"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetJWTConfiguration(name string) string {
	if name == "exp" {
		conf := configs.GetConfigJWT()
		return conf.JWT_EXP
	}
	return ""
}
