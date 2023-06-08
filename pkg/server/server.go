package server

import (
	"challenge/pkg/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	routes.RegisterApiRoutes(app)

	return app
}
