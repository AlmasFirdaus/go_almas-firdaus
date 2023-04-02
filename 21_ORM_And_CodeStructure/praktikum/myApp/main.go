package main

import (
	"myApp/config"
	"myApp/routes"
)

func main() {
	config.InitDB()
	// start the server, and log if it fails
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
