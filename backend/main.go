package main

import (
	"gofiber-restful-api/app"
	"gofiber-restful-api/database"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.InitDB()

	fiberApp := fiber.New()
	app.SetupRouter(fiberApp)

	log.Fatal(fiberApp.Listen(":8000"))
}