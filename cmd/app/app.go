package app

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ologbonowiwi/toggl-challenge/internal/api"
	"github.com/ologbonowiwi/toggl-challenge/internal/services"
	"github.com/ologbonowiwi/toggl-challenge/internal/storage"
)

func SetupApp() *fiber.App {
	deckRepository := storage.NewLocalDeckRepository()

	service := services.NewDeckService(deckRepository)

	app := fiber.New()

	apiHandler := api.NewAPIHandler(*service)

	apiHandler.SetupRoutes(app.Group("/api"))

	return app
}
