package main

import (
	"crud/webServer"
	"crud/modelResource"
)

func main() {
	// connect database
	// persistResource.ConnectDatabase()

	//	next line just generates some test data records:
	modelResource.Init()

	// start webserver
	crud.StartWebserver()
}
