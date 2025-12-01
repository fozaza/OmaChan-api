package router

import (
	"fmt"

	"github.com/OmaChan/database/table"
	"github.com/gofiber/fiber/v2"
)

// auto create hardware user
func acr_hw(c *fiber.Ctx) error {
	var req table.HardWareInput
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	if err := table.Cr_ha(req); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	hardware, err := table.Ge_ha(req.MapHard().Name)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := table.Cr_hwd(hardware.MapHard(), ""); err != nil {
		return err
	}
	return c.SendString("Success")
}

// man
