package main

import (
	"log"
	"os"

	"github.com/ologbonowiwi/toggl-challenge/cmd/app"
)

func main() {
	app := app.SetupApp()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
