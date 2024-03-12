package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peammapake/vincent-inventory-api/internal/server"
)

func NotFoundRoute(a *fiber.App, s *server.Server){
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg": "endpoint not found",
			})
		},
	)
}