package main

import (
	"altera/Day5-6/pkg/config"
	"altera/Day5-6/pkg/helper"
	"altera/Day5-6/pkg/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	helper.LogMiddleware(e)
	e.Start(":8080")
}
