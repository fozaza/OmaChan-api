package router

import (
	"github.com/OmaChan/database/table"
	"github.com/OmaChan/module"
	"github.com/gofiber/fiber/v2"
)

func create_user(c *fiber.Ctx) error {
	// read input from json
	var user_input table.UserInput
	if err := c.BodyParser(&user_input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("OmaChan >>> BadRequest plz input name email and password")
	}

	// create to database
	if err := table.Cr_user(user_input); err.Err != nil {
		return err.MapFiber(c)
	}
	return c.SendString("OmaChan >>> Success to create id")
}

func login_user(c *fiber.Ctx) error {
	// read input from json
	var req table.UserLogin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	// check users and password
	user, err := table.Login(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> Error to login")
	}

	t, err := module.Cr_jwt(user.Email, user.Level_user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"token": t,
		"Msg":   "Hello",
	})
}

func Get_all_router(r fiber.Router) {
	r.Post("/login", login_user)
	r.Post("/Create", create_user)
}
