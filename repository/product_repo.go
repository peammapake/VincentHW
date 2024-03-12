package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peammapake/vincent-inventory-api/model/product"
	"github.com/peammapake/vincent-inventory-api/pkg/pagination"
	"github.com/peammapake/vincent-inventory-api/utils/common"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductList(c *fiber.Ctx, filter GetProductListFilter) ([]product.Products, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

type GetProductListFilter struct {
	ProductCode *string
	ProductName *string
	BrandId     *string
	CategoryId  *string
	MinNetPrice *int32
	MaxNetPrice *int32
	InStock     *bool
	Pagination  common.PaginationRequest
}

func (r *ProductRepositoryImpl) GetProductList(c *fiber.Ctx, filter GetProductListFilter) ([]product.Products, error) {
	var products []product.Products

	query := pagination.PaginationAndSort(r.db, filter.Pagination)

	if filter.ProductCode != nil {
		query = query.Where("REGEXP_LIKE(code, ?, 'i')", filter.ProductCode)
	}

	if filter.ProductName != nil {
		query = query.Where("REGEXP_LIKE(name, ?, 'i')", filter.ProductName)
	}

	if filter.BrandId != nil {
		query = query.Where("brand")
	}

	result := query.Find(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}
