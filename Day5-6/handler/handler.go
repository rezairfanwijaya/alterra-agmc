package handler

import (
	"altera/Day5-6/pkg/helper"
	"altera/Day5-6/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (uh *userHandler) GetAllUser(e echo.Context) error {
	user, err := uh.service.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (uh *userHandler) GetUserById(e echo.Context) error {
	// id validation
	userID, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := uh.service.GetUserById(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (uh *userHandler) DeleteUserById(e echo.Context) error {
	// id validation
	userID, err := helper.IdValidator(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = uh.service.DeleteUserById(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, "successfully deleted user")
}
