package router

import (
	"github.com/williamroberttv/curriculum-gen-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//setup the user group
	user := app.Group("/user")
	user.Get("/:id", services.GetUser)
	user.Post("/", services.CreateUser)

	// setup theauth group
	auth := app.Group("/auth")
	auth.Post("/", services.Authenticate)
	auth.Get("/", services.RefreshToken)
}