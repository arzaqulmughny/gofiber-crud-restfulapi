package main

import (
	"gofiber-restful-api/app"
	"gofiber-restful-api/database"
	"log"

	_ "gofiber-restful-api/docs"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
)

// @title			Notes API
// @version		1.0
// @description	RESTful API documentation for notes api
// @BasePath	/api
func main() {
	godotenv.Load()
	database.InitDB()

	fiberApp := fiber.New()
	app.SetupRouter(fiberApp)

	fiberApp.Get("/swagger/*", swaggo.HandlerDefault)

	log.Fatal(fiberApp.Listen(":8000"))
}
