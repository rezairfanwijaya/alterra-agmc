package routes

import (
	"altera/Day5-6/handler"
	"altera/Day5-6/pkg/config"
	"altera/Day5-6/user"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// user repo
	userRepo := user.NewRepository(config.DB)
	// user service
	userService := user.NewService(userRepo)
	// user handler
	userHandler := handler.NewUserHandler(userService)

	// endpoint
	api := e.Group("api/v1")
	api.GET("/users", userHandler.GetAllUser)
	api.GET("/user/:id", userHandler.GetUserById)

	return e
}
