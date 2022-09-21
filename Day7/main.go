package main

import (
	"altera/Day7/pkg/config"
	"altera/Day7/pkg/helper"
	"altera/Day7/pkg/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	helper.LogMiddleware(e)
	e.Start(":8080")
}
