package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type Categories struct {
	Id        *uuid.UUID `db:"id"`
	Name      string     `db:"name"`
	DeletedAt *time.Time `db:"deleted_at"`
	model.Auditable
}
