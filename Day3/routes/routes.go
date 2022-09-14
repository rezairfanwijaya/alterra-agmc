package routes

import (
	"altera/Day3/config"
	"altera/Day3/controllers"
	"altera/Day3/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
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
	e.POST("/login", controllers.LoginUser)
	e.PUT("/user/:id", controllers.UpdateUserById)

	// implementasi middleware basic auth
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middleware.BasicAuthDB))

	eAuth.DELETE("/user/:id", controllers.DeleteUserById)

	// implementasi middlewate jwt
	jwt := e.Group("/jwt")
	data := config.LoadENV()
	jwt.Use(m.JWT([]byte(data["jwtSecret"])))

	jwt.GET("/user/detail", controllers.GetUserById)

	return e
}
