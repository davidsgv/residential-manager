package api

import (
	"residential-manager/internal/common/response"

	"github.com/gofiber/fiber/v2"
)

func (app *Aplication) loginHandler(c *fiber.Ctx) error {
	var input struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	srvResponse := app.Services.AuthService.Login(input.Mail, input.Password)
	apiResponse := response.Response[response.M]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
	}
	if srvResponse.Data != nil {
		apiResponse.Data = &response.M{"token": srvResponse.Data}
	}

	//return c.JSON(srvResponse)
	return app.restResponse(c, apiResponse.Status, apiResponse)
}
