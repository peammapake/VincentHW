package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	exception "github.com/peammapake/vincent-inventory-api/pkg/exeption"
	"github.com/peammapake/vincent-inventory-api/pkg/validators"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=50"`
	LastName  string `json:"lastName" validate:"required,min=1,max=50"`
	Username  string `json:"username" validate:"required,min=5,max=30"`
	Password  string `json:"password" validate:"required,min=8,max=16"`
	RoleId    string `json:"roleId" validate:"required"`
}

type CreateUserResponse struct {
	Id     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (s *Server) CreateUser(c *fiber.Ctx) error {
	// Parse JSON request from context
	var request GetStockProductListRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exception.NewValidationError(err))
	}

	// Validate request field validation
	if err := validators.ValidateStruct(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exception.NewValidationError(err))
	}

	return c.JSON(nil)
}
