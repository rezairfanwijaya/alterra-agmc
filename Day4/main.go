package main

import (
	"altera/Day2/config"
	"altera/Day3/routes"
	m "altera/Day3/middleware"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Start(":8080")
}
