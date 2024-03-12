package user

import (
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type RolePermissions struct {
	RoleId     uuid.UUID `db:"role_id"`
	Permission string    `db:"permission"`
	model.Auditable
}
