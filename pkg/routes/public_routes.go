package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peammapake/vincent-inventory-api/internal/server"
)

func PublicRoutes(a *fiber.App, s *server.Server) {
	// Create routes group
	routes := a.Group("/api")

	// GET
	routes.Get("/items", s.GetStockProductList)
}
