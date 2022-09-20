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

func (uh *userHandler) AddNewUser(e echo.Context) error {
	var user user.UserInput

	// binding
	if err := e.Bind(&user); err != nil {
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, err)
		return e.JSON(http.StatusBadRequest, response)
	}

	// validate
	if err := e.Validate(&user); err != nil {
		myErr := helper.ErrorBind(err)
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, myErr)
		return e.JSON(http.StatusBadRequest, response)
	}

	// panggil service
	newUser, err := uh.service.AddNewUser(user)
	if err != nil {
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI("success", "success", http.StatusOK, newUser)
	return e.JSON(http.StatusOK, response)

}

func (uh *userHandler) UpdateUserById(e echo.Context) error {
	// id validation
	userID, err := helper.IdValidator(e)
	if err != nil {
		response := helper.ResponseAPI(err.Error(), "failed", http.StatusBadRequest, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	var userInput user.UserInput

	// binding
	if err := e.Bind(&userInput); err != nil {
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, err)
		return e.JSON(http.StatusBadRequest, response)
	}

	// validate
	if err := e.Validate(&userInput); err != nil {
		myErr := helper.ErrorBind(err)
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, myErr)
		return e.JSON(http.StatusBadRequest, response)
	}

	// panggil service
	userUpdated, err := uh.service.UpdateUserById(userInput, userID)
	if err != nil {
		response := helper.ResponseAPI("failed", "failed", http.StatusBadRequest, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	return e.JSON(http.StatusBadRequest, userUpdated)
}
