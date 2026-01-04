package router

import (
	// "fmt"

	// "github.com/OmaChan/database/table"

	"github.com/OmaChan/database/table"
	"github.com/gofiber/fiber/v2"
)

// auto create hardware user

type HardWareInput struct {
	Name  string
	Title string
}

func acr_hw(c *fiber.Ctx) error {
	var req HardWareInput
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	// map data
	var hardwareNew table.HardWare
	hardwareNew.Name = req.Name
	hardwareNew.Title = req.Title
	hardwareNew.Enable = true

	err := table.Cr_ha(hardwareNew) // Create data
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("OmaChan >>> Error " + err.Error())
	}

	return c.SendString("Success")
}
func up_hw(c *fiber.Ctx) error {
	type Data struct {
		Name   string
		Pm     float32
		Batter float32
	}

	data := table.Data{
		Pm:     float32(c.QueryFloat("Pm")),
		Batter: float32(c.QueryInt("Batter")),
	}

	// println(float32(c.QueryFloat("Pm")))
	// println(float32(c.QueryInt("Batter")))
	if err := table.Up_data(data, c.Query("Name")); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.SendString("OmaChan >>> Update hardware data")
}

// man
