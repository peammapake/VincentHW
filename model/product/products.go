package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model"
	"github.com/shopspring/decimal"
)

type Products struct {
	Id                   *uuid.UUID       `db:"id"`
	Name                 string           `db:"name"`
	Description          *string          `db:"description"`
	HighlightCoverImgUrl *string          `db:"highlight_cover_img_url"`
	BrandId              *uuid.UUID       `db:"brand_id"`
	CategoryId           *uuid.UUID       `db:"category_id"`
	Code                 string           `db:"code"`
	NetPrice             *decimal.Decimal `db:"net_price"`
	GrossPrice           *decimal.Decimal `db:"gross_price"`
	Stock                int32            `db:"stock"`
	DeletedAt            *time.Time       `db:"deleted_at"`
	model.Auditable
}
