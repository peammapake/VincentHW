package user

import (
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type UserRoles struct {
	Id       *uuid.UUID `db:"id"`
	RoleName string     `db:"role_name"`
	model.Auditable
}
