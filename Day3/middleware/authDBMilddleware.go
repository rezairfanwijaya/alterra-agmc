package middleware

import (
	"altera/Day2/models"
	"altera/Day3/config"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	user := models.User{}
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}
