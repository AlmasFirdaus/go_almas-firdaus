package main

import (
	"myAppMiddleware/config"
	m "myAppMiddleware/middlewares"
	"myAppMiddleware/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	// implement middleware logger
	m.LogMiddlewares(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
