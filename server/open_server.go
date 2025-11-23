package server

import (
	"github.com/OmaChan/module"
	"github.com/gofiber/fiber/v2"
)

func OpenServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello For OmaChan")
	})

	module.Con_jwt(app)
	module.ExtractUserFromJWT(app)

	admin := app.Group("/admin")
	admin.Use(module.Req_level(4))

	app.Listen("0.0.0.0:3000")
}
