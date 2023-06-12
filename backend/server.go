package main

import (
	"tandigital/backend/db"
	"tandigital/backend/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// app.Get("/admins", handlers.GetAllAdmins)
	router.SetupRouter(app)

	app.Listen(":8080")

}
