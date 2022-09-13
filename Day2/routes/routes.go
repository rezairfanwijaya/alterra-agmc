package routes

import (
	"altera/Day2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.AddUser)

	return e

}
