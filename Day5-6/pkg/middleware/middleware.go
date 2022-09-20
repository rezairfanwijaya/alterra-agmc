package middleware

import (
	"altera/Day5-6/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int, e echo.Context) (string, error) {
	claims := jwt.MapClaims{
		"authorization": true,
		"userId":        userId,
		"exp":           time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	data := config.LoadENV()
	secret := data["jwtSecret"]

	return token.SignedString([]byte(secret))
}
