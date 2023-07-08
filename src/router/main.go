package router

import (
	"github.com/williamroberttv/curriculum-gen-api/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// setup the auth group
	auth := app.Group("/auth")
	auth.Get("/", handlers.Authenticate)
}