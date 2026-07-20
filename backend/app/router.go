package app

import (
	"gofiber-restful-api/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRouter(app *fiber.App) {
	// Group API
	api := app.Group("/api")

	// Notes
	notes := api.Group("/notes")
	notes.Post("/", controllers.CreateNotesController)
	notes.Get("/", controllers.FindAllNotesController)
}