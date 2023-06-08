package routes

import "github.com/gofiber/fiber/v2"

func RegisterApiRoutes(router fiber.Router) {
	api := router.Group("/api")

	RegisterJokesRoutes(api)
}
