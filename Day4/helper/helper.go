package helper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorBind(err error) []string {
	var myErr []string
	for _, e := range err.(validator.ValidationErrors) {
		// generate error from validator
		errMessage := fmt.Sprintf("error on filed: %v, condition: %v", e.Field(), e.ActualTag())
		myErr = append(myErr, errMessage)
	}
	return myErr
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
