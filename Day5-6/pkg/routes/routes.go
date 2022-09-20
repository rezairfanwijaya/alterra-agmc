package routes

import (
	"altera/Day5-6/handler"
	"altera/Day5-6/pkg/config"
	"altera/Day5-6/user"

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

	// user repo
	userRepo := user.NewRepository(config.DB)
	// user service
	userService := user.NewService(userRepo)
	// user handler
	userHandler := handler.NewUserHandler(userService)

	// endpoint
	api := e.Group("api/v1")
	api.GET("/users", userHandler.GetAllUser)
	// api.GET("/user/:id", userHandler.GetUserById)
	api.DELETE("/user/:id", userHandler.DeleteUserById)
	api.POST("/user", userHandler.AddNewUser)
	api.PUT("/user/:id", userHandler.UpdateUserById)
	api.POST("/auth", userHandler.Auth)

	// implementasi middlewate jwt
	jwt := e.Group("api/v1/jwt")
	data := config.LoadENV()
	jwt.Use(m.JWT([]byte(data["jwtSecret"])))

	jwt.GET("/user/:id", userHandler.GetUserById)

	return e
}
