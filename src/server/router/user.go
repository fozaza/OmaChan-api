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
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input name email and password")
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

	t, err := module.Cr_jwt(user.Email, user.Level)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"token": t,
		"Msg":   "Hello",
	})
}

func get_user(c *fiber.Ctx) error {
	var req table.QueryUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	user, err := table.Gt_user(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"Msg":  "Req ok",
		"User": user,
	})
}

func get_user_all(c *fiber.Ctx) error {
	var req table.QueryUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	user_all, err := table.Gt_all_user(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"Msg":  "Req Ok",
		"User": user_all,
	})
}

func rm_self(c *fiber.Ctx) error {
	// read rm yes_no
	var password table.Password

	if err := c.BodyParser(&password); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	email, _ := module.Get_token(c)
	user := table.UserLogin{
		Email:    email,
		Password: string(password),
	}

	err := table.Rm_self(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.SendString("OmaChan >> remove account success")
}

func change_level(c *fiber.Ctx) error {
	// read input json
	var req table.UserRetrun
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	// Get email req and user
	var admin table.UserRetrun
	admin.Email, admin.Level = module.Get_token(c)

	if result := table.Ch_le(admin, req.Email, req.Level); result != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(result.Error())
	}
	return c.SendString("OmaChan >>> update level user success")
}

func remove_user(c *fiber.Ctx) error {
	// read josn
	var req table.RemoveUserWithAdmin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> BadRequest plz input email and password")
	}

	result, err := table.Rm_user(req.Admin, req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"Massage": "OmaChan remove user list",
		"Log":     result,
	})
}

func Get_all_router(r fiber.Router) {
	r.Post("/login", login_user)
	r.Post("/create", create_user)
	r.Post("/remove", rm_self)
	r.Post("/User", get_user)
	r.Post("/AllUser", get_user_all)
	r.Post("/Gt_hwd", gt_hwd)

	r.Post("/admin/changeLevel", change_level)
	r.Post("/admin/removeUser", remove_user)
	r.Post("/admin/createHwd", acr_hw)
}
