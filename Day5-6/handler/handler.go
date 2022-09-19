package handler

import (
	"altera/Day5/user"
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
