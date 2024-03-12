package repository

import (
	"github.com/peammapake/vincent-inventory-api/initializer"
	"gorm.io/gorm"
)

type Repository struct {
	dbPool *gorm.DB

	Product ProductRepository
}

func NewRepository() *Repository {
	dbPool := initializer.Db
	return &Repository{
		dbPool:  dbPool,
		Product: NewProductRepository(dbPool),
	}
}
