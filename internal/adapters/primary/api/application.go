package api

import (
	"residential-manager/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type Services struct {
	AuthService service.AuthSrv
	UserService service.UserSrv
}

type Aplication struct {
	Fiber  *fiber.App
	config *Config
	Services
}

func NewAplication(cfg *Config, services Services) *Aplication {
	app := Aplication{
		Fiber:    fiber.New(),
		config:   cfg,
		Services: services,
	}
	app.routes()

	return &app
}
