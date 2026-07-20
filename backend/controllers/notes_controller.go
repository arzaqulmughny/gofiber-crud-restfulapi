package controllers

import (
	"gofiber-restful-api/domain"
	"gofiber-restful-api/helpers"
	"gofiber-restful-api/services"

	"github.com/gofiber/fiber/v3"
)

func CreateNotesController(c fiber.Ctx) error {
	notes := new(domain.Notes)

	// Parse request body (JSON) to notes struct
	if err := c.Bind().JSON(notes); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if errors := helpers.Validate(notes); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	// Try store new notes
	id, err := services.CreateNotesService(notes)

	if err != nil {
		// Failed to store new notes
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Success store new notes
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func FindAllNotesController (c fiber.Ctx) error {
	notes, err := services.FindAllNotesSerivice()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": notes,
		"error": nil,
	})
}