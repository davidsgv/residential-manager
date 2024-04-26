package api

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (app *Aplication) routes() {
	//app.Fiber.Use(app.coorMiddleware)
	app.Fiber.Use(cors.New())

	app.Fiber.Get("/v1/healthcheck", app.healthcheckHandler)
	app.Fiber.Post("/v1/auth/login", app.loginHandler)
	app.Fiber.Get("/v1/auth/verify/:token", app.showVerifyHandler)
	app.Fiber.Post("/v1/auth/verify/:token", app.verifyHandler)

	//middleware
	apiSecure := app.Fiber.Group("/v1", app.authMiddleware)

	apiSecure.Get("/permissions", app.showPermissions)
	apiSecure.Get("/roles", app.showRoles)

	apiSecure.Post("/users", app.createUserHandler)
	apiSecure.Get("/users", app.showUsersHandler)
	apiSecure.Get("/users/:id", app.showUserHandler)
	apiSecure.Put("/users/:id", app.updateUserHandler)
	apiSecure.Delete("/users/:id", app.deleteUserHandler)
}
