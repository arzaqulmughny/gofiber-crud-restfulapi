package controllers

import (
	"gofiber-restful-api/dtos"
	"gofiber-restful-api/helpers"
	"gofiber-restful-api/services"

	"github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
	var request dtos.UserLoginRequest

	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validate request
	if err := helpers.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	token, err := services.UserLoginService(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login succesful",
		"data":    token,
	})
}

func Register(c fiber.Ctx) error {
	var request dtos.UserRegisterRequest

	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	errorValidation := helpers.Validate(request)

	if errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidation,
		})
	}

	user, err := services.UserRegisterService(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create user successful",
		"data":    user,
	})
}
