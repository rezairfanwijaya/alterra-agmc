package main

import (
	"altera/Day2/config"
	"altera/Day2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}
