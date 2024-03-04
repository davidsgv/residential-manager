package api

import (
	"github.com/gofiber/fiber/v2"
)

func (app *Aplication) healthcheckHandler(c *fiber.Ctx) error {
	data := map[string]string{
		"status":      "avalaible",
		"environment": app.config.Env,
		"version":     version,
	}

	return c.JSON(data)
}
