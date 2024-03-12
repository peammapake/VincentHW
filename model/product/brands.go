package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type Brands struct {
	Id           *uuid.UUID `db:"id"`
	Name         string     `db:"name"`
	Code         *string    `db:"code"`
	BrandIconUrl *string    `db:"brand_icon_url"`
	DeletedAt    *time.Time `db:"deleted_at"`
	model.Auditable
}
