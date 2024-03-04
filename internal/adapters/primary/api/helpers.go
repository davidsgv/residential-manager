package api

import (
	"residential-manager/internal/common/response"
	srvrequest "residential-manager/internal/common/srvRequest"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// func (app *Aplication) readIDParam(c *fiber.Ctx) (int64, error) {
// 	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
// 	if err != nil || id < 1 {
// 		return id, fiber.NewError(fiber.StatusBadRequest, "invalid id parameter")
// 	}
// 	return id, nil
// }

func (app *Aplication) readUUIDParam(c *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return id, fiber.NewError(fiber.StatusBadRequest, "invalid id parameter")
	}
	return id, nil
}

func (app *Aplication) restResponse(c *fiber.Ctx, status string, data any) error {
	if status == response.StatusSuccess {
		return c.Status(fiber.StatusAccepted).JSON(data)
	}
	if status == response.StatusFail {
		return c.Status(fiber.StatusBadRequest).JSON(data)
	}
	if status == response.StatusUnauthorized {
		return c.Status(fiber.StatusUnauthorized).JSON(data)
	}
	if status == response.StatusNotFound {
		return c.Status(fiber.StatusNotFound).JSON(data)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(data)
}

func (app *Aplication) getServiceRequest(c *fiber.Ctx) srvrequest.Request {
	return c.Locals("srv").(srvrequest.Request)
}
