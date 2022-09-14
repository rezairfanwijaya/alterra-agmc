package routes

import (
	"altera/Day2/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.AddUser)
	e.GET("/user/:id", controllers.GetUserById)
	e.DELETE("/user/:id", controllers.DeleteUserById)

	return e

}
