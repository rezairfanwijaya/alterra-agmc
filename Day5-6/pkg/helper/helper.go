package helper

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func IdValidator(e echo.Context) (int, error) {
	// get param
	param := e.Param("id")

	// convert to int
	id, err := strconv.Atoi(param)
	if err != nil {
		return id, errors.New("id must be integer and grater than 0")

	}

	// validasi
	if id <= 0 {
		return id, errors.New("id must be integer and grater than 0")
	}

	return id, nil
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
}