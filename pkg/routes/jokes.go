package routes

import (
	"challenge/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterJokesRoutes(router fiber.Router) {
	api := router.Group("/jokes")

	api.Get("/", handlers.GetJokes)
}
