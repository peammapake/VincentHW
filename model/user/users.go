package user

import (
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type Users struct {
	Id        *uuid.UUID `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	RoleId    uuid.UUID  `db:"role_id"`
	model.Auditable
}
