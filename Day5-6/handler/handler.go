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
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI("Success", "Success", http.StatusOK, user)
	return e.JSON(http.StatusOK, response)
}

func (uh *userHandler) GetUserById(e echo.Context) error {
	// id validation
	userID, err := helper.IdValidator(e)
	if err != nil {
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	user, err := uh.service.GetUserById(userID)
	if err != nil {
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI("Success", "Success", http.StatusOK, user)
	return e.JSON(http.StatusOK, response)
}

func (uh *userHandler) DeleteUserById(e echo.Context) error {
	// id validation
	userID, err := helper.IdValidator(e)
	if err != nil {
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = uh.service.DeleteUserById(userID)
	if err != nil {
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI("Success", "Success", http.StatusOK, nil)
	return e.JSON(http.StatusOK, response)
}
