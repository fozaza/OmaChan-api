package server

import "github.com/gofiber/fiber/v2"

func OpenServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello For OmaChan")
	})

	app.Listen("0.0.0.0:3000")
}
