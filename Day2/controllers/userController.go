package controllers

import (
	"altera/Day2/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(e echo.Context) error {
	// panggil database user
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}
