package router

import (
	"tandigital/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	admin := app.Group("/admin")
	admin.Get("/", handlers.GetAllAdmins)
	admin.Post("/", handlers.CreateAdmin)
	admin.Post("/test", handlers.TestBody)
}
