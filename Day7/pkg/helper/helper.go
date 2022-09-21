package helper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

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

func ResponseAPI(message, status string, code int, data interface{}) *response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	return &response{
		Meta: meta,
		Data: data,
	}
}

func ErrorBind(err error) []string {
	var myErr []string
	for _, e := range err.(validator.ValidationErrors) {
		// generate error from validator
		errMessage := fmt.Sprintf("error on filed: %v, condition: %v", e.Field(), e.ActualTag())
		myErr = append(myErr, errMessage)
	}
	return myErr
}
