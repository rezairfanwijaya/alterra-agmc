package controllers

import (
	"altera/Day2/lib/database"
	"altera/Day2/models"
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

func AddUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}
