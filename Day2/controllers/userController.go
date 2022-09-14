package controllers

import (
	"altera/Day2/helper"
	"altera/Day2/lib/database"
	"altera/Day2/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all user
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

// add user
func AddUser(e echo.Context) error {
	var user models.User

	// binding
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate
	if err := e.Validate(&user); err != nil {
		errBind := helper.ErrorBind(err)
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": errBind,
		})
	}

	// cek apakah user sudah terdaftar ?
	data, err := database.FindUserByEmail(user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// jika email sama maka user sudah regis dan munculkan error
	if data.Email == user.Email {
		msg := fmt.Sprintf("user dengan email %v sudah terdaftar", user.Email)
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": msg,
		})
	}

	// add user jika user belum terdaftar
	_, err = database.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully add user",
	})

}

// get user by id
func GetUserById(e echo.Context) error {
	// id validation
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// find user by id
	user, err := database.FindUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user.Email == "" {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "user not found",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully get user by id",
		"data":    user,
	})

}

// update user by id
func UpdateUserById(e echo.Context) error {
	// id validation
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// get user by id
	user, err := database.FindUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// binding
	var inputUser models.User

	if err := e.Bind(&inputUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate
	if err := e.Validate(&inputUser); err != nil {
		errBind := helper.ErrorBind(err)
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": errBind,
		})
	}

	user.Email = inputUser.Email
	user.Name = inputUser.Name
	user.Password = inputUser.Password

	// update
	userUpdated, err := database.UpdateUserById(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update user",
		"data":    userUpdated,
	})

}

// delete user by id
func DeleteUserById(e echo.Context) error {
	// id validation
	id, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// delete
	if err = database.DeleteUserById(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully delete user",
	})
	

}
