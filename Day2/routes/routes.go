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

	// dinamis
	e.GET("/users", controllers.GetUsers)
	e.GET("/user/:id", controllers.GetUserById)
	e.POST("/user", controllers.AddUser)
	e.PUT("/user/:id", controllers.UpdateUserById)
	e.DELETE("/user/:id", controllers.DeleteUserById)

	// statis
	e.GET("/books", controllers.GetAllBook)
	e.GET("/books/:id", controllers.GetBookById)
	e.PUT("/books/:id", controllers.UpdateBookById)
	e.DELETE("/books/:id", controllers.DeleteBookById)
	e.POST("/books", controllers.AddBook)

	return e

}
