package main

import (
	"log"

	"github.com/ologbonowiwi/toggl-challenge/cmd/app"
)

func main() {
	app := app.SetupApp()

	log.Fatal(app.Listen(":3000"))
}
