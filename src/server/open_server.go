package server

import (
	"strconv"

	"github.com/OmaChan/module"
	"github.com/OmaChan/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func OpenServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello For OmaChan")
	})

	api := app.Group("/")
	router.Get_all_router(api)

	module.Con_jwt(app)
	//module.ExtractUserFromJWT(app)
	api.Get("/jwtCheck", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		println(claims)
		name := claims["email"].(string)
		level := claims["level"].(float64)
		//println(int(level))
		return c.SendString("Welcome >>> " + name + ": level >>> " + strconv.FormatFloat(level, 'f', 0, 64))
	})

	admin := app.Group("/admin")
	admin.Use(module.Req_level(4))
	admin.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello u r admin")
	})

	app.Listen("0.0.0.0:3000")
}
