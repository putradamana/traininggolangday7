package main

import (
	"go-training-restful/routes"
)

func main() {
	e := routes.New()

	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
