package api

import (
	"residential-manager/internal/common/response"
	srvrequest "residential-manager/internal/common/srvRequest"

	"github.com/gofiber/fiber/v2"
)

func (app *Aplication) authMiddleware(c *fiber.Ctx) error {
	clientToken := c.Get("Authorization")
	data := app.Services.AuthService.CheckLogin(clientToken)
	if data.Error != nil {
		res := response.Response[any]{}
		res = res.ResponseUnauthorized(data.Error.Error())
		return app.restResponse(c, res.Status, res)
	}

	c.Locals("srv", srvrequest.Request{
		Mail: data.Mail,
	})

	return c.Next()
}

func (app *Aplication) coorMiddleware(c *fiber.Ctx) error {
	c.Append("Access-Control-Allow-Origin", "*")
	return c.Next()
}
