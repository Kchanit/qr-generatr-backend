package router

import (
	"github.com/Kchanit/qr-generator/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	api.Post("/generate", controllers.HandleGenerateQRCode)
}
