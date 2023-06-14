package main

import (
	"tandigital/backend/database"
	"tandigital/backend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.Init()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// app.Get("/admins", handlers.GetAllAdmins)
	router.SetupRouter(app)

	app.Listen(":8080")

}
