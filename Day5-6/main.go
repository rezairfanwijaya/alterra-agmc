package main

import (
	"altera/Day5/pkg/config"
	"altera/Day5/pkg/helper"
	"altera/Day5/pkg/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	helper.LogMiddleware(e)
	e.Start(":8080")

}
