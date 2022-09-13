package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello"})
}

func user(c echo.Context) error {
	users := []User{
		{
			Id:      1,
			Name:    "John",
			Address: "New York",
		},
		{
			Id:      2,
			Name:    "Doe",
			Address: "Singapure",
		},
		{
			Id:      3,
			Name:    "Peter",
			Address: "Malaysia",
		},
		{
			Id:      4,
			Name:    "Parker",
			Address: "California",
		},
		{
			Id:      5,
			Name:    "Amelia",
			Address: "Jonggol",
		},
	}

	return c.JSON(http.StatusOK, users)
}

func employee(c echo.Context) error {
	employees := []User{
		{
			Id:      1000,
			Name:    "Anwar",
			Address: "Jakarta",
		},
		{
			Id:      1001,
			Name:    "Joko",
			Address: "Surabaya",
		},
		{
			Id:      1002,
			Name:    "Widodo",
			Address: "Purwokerto",
		},
	}

	return c.JSON(http.StatusOK, employees)
}

func main() {
	// init echo
	e := echo.New()

	// create route
	e.GET("/home", home)
	e.GET("/user", user)

	// grouping
	v1 := e.Group("v1")
	v1.GET("/employee", employee)

	// run server
	e.Start(":7070")
}
