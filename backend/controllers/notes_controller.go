package controllers

import (
	"gofiber-restful-api/domain"
	"gofiber-restful-api/dtos"
	"gofiber-restful-api/helpers"
	"gofiber-restful-api/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// new note godoc
//
// @Summary Store new note
// @Description Create new note
// @Tags Notes
// @Accept json
// @Param body body dtos.CreateNewNoteRequest true "Data"
// @Produce json
// @Success 201 {object} dtos.CreateNewNoteResponse
// @Router /notes [post]
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

	// Return the created note
	note, err := services.FindNoteByIdService(uint(id))

	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Data note tidak ditemukan.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": nil,
		"data":  note,
	})
}

// Get all notes godoc
//
// @Summary Show all notes
// @Description Get all notes
// @Tags Notes
// @Accept json
// @Param body body dtos.CreateNewNoteRequest true "Data"
// @Produce json
// @Success 200 {object} dtos.FindAllNotesResponse
// @Router /notes [get]
func FindAllNotesController(c fiber.Ctx) error {
	notes, err := services.FindAllNotesSerivice()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  notes,
		"error": nil,
	})
}

// Find note godoc
//
// @Tags Notes
// @Param id path int true "Note id"
// @Produce json
// @Success 200 {object} dtos.CreateNewNoteResponse
// @Summary Find note by id
// @Router /notes/{id} [get]
func FindNoteByIdController(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Id tidak valid.",
		})
	}

	// Find note by id
	note, err := services.FindNoteByIdService(uint(id))

	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Data note tidak ditemukan.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": nil,
		"data":  note,
	})
}

func DeleteNoteByIdController(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Id tidak valid.",
		})
	}

	err = services.DeleteNoteByIdService(uint(id))

	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   nil,
		"message": "Note berhasil dihapus.",
	})
}

// Update note godoc
//
// @Summary Update existing note
// @Tags Notes
// @Param id path int true "Note id"
// @Produce json
// @Success 200 {object} dtos.CreateNewNoteResponse
// @Router /notes/{id} [put]
func UpdateNoteByIdController(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	var request dtos.UpdateNoteRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	result, err := services.UpdateNoteByIdService(uint(id), &request)

	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
		"message": "Data update successful",
	})
}
