package product

import (
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
)

type ProductImages struct {
	ProductId uuid.UUID `db:"product_id"`
	ImageUrl  string    `db:"image_url"`
	model.Auditable
}
