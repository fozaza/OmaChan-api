package router

import (
	"github.com/OmaChan/database/table"
	"github.com/gofiber/fiber/v2"
)

func gt_hwd(c *fiber.Ctx) error {
	var req string
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	hwd, err := table.Ge_ha(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}
	return c.JSON(fiber.Map{"Data": hwd})
}
