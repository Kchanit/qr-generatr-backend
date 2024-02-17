package main

import (
	"github.com/Kchanit/qr-generator/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "https://qr-generatr.vercel.app",
			AllowCredentials: true,
		}))

	router.SetUpRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
