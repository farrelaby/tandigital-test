package router

import (
	"tandigital/backend/handlers"
	"tandigital/backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	admin := app.Group("/admin")
	admin.Get("/list", handlers.GetAllAdmins)
	admin.Post("/register", handlers.CreateAdmin)
	admin.Post("/login", handlers.LoginAdmin)
	// admin.Post("/test", handlers.TestBody)

	voucher := app.Group("/voucher")
	voucher.Get("/list", handlers.GetAllVouchers)

	manage := app.Group("/manage")
	manage.Use(middleware.CheckAdminToken)
	manage.Post("/create", handlers.CreateVoucher)
	manage.Delete("/delete/:id", handlers.DeleteVoucher)

	consumer := app.Group("/consumer")
	consumer.Post("/register", handlers.RegisterConsumer)
	consumer.Post("/login", handlers.LoginConsumer)

	transaction := app.Group("/transaction")
	transaction.Use(middleware.CheckConsumerToken)
	transaction.Post("/buy/:id", handlers.BuyVoucher)
	transaction.Get("/list", handlers.GetVoucherbyConsumer)

}
