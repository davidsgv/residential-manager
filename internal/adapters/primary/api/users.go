package api

import (
	"residential-manager/internal/adapters/primary/api/data"
	"residential-manager/internal/common/response"
	"residential-manager/internal/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (app *Aplication) showPermissions(c *fiber.Ctx) error {
	srvResponse := app.Services.UserService.GetAllPermissions(app.getServiceRequest(c))
	permissions := data.MapToPermissions(*srvResponse.Data)
	apires := response.Response[response.M]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        &response.M{"permissions": &permissions},
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) showRoles(c *fiber.Ctx) error {
	srvResponse := app.Services.UserService.GetRoles(app.getServiceRequest(c))
	apires := response.Response[response.M]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        &response.M{"roles": *srvResponse.Data},
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) createUserHandler(c *fiber.Ctx) error {
	var input data.CreateUser
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	srvResponse := app.Services.UserService.CreateUser(app.getServiceRequest(c), input.Rol, input.Mail, input.Apartment.Block, input.Apartment.Number)
	apires := response.Response[data.UserResponse]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        data.UserToApi(srvResponse.Data),
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) showUsersHandler(c *fiber.Ctx) error {
	srvResponse := app.Services.UserService.GetUsers(app.getServiceRequest(c))
	users := data.UsersToApi(*srvResponse.Data)
	apires := response.Response[[]data.UserResponse]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        &users,
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) showUserHandler(c *fiber.Ctx) error {
	id, err := app.readUUIDParam(c)
	if err != nil {
		return err
	}

	srvResponse := app.Services.UserService.GetUserById(app.getServiceRequest(c), id)

	apires := response.Response[data.UserResponse]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        data.UserToApi(srvResponse.Data),
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) updateUserHandler(c *fiber.Ctx) error {
	var input data.UpdateUser
	id, err := app.readUUIDParam(c)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	user := entities.User{
		Id:        id,
		Rol:       input.Rol,
		Apartment: nil,
	}
	if input.Apartment != nil {
		user.Apartment = &entities.Apartment{
			Block:  input.Apartment.Block,
			Number: input.Apartment.Number,
		}
	}

	srvResponse := app.Services.UserService.UpdateUser(app.getServiceRequest(c), user)
	apires := response.Response[data.UserResponse]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        data.UserToApi(srvResponse.Data),
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) deleteUserHandler(c *fiber.Ctx) error {
	id, err := app.readUUIDParam(c)
	if err != nil {
		return err
	}

	srvResponse := app.Services.UserService.DeleteUser(app.getServiceRequest(c), id)
	apires := response.Response[map[string]string]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        nil,
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) showVerifyHandler(c *fiber.Ctx) error {
	token := c.Params("token")
	srvResponse := app.Services.UserService.GetVerifyToken(token)
	apires := response.Response[data.UserResponse]{
		Status:      srvResponse.Status,
		Message:     srvResponse.Message,
		Validations: srvResponse.Validations,
		Data:        data.UserToApi(srvResponse.Data),
	}
	return app.restResponse(c, srvResponse.Status, apires)
}

func (app *Aplication) verifyHandler(c *fiber.Ctx) error {
	token := c.Params("token")

	var input data.VerifyUser
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	srvResponse := app.Services.UserService.VerifyToken(token, input.Password)
	return app.restResponse(c, srvResponse.Status, srvResponse)
}
