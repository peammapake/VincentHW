package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/model/product"
	"github.com/peammapake/vincent-inventory-api/repository"
	"github.com/peammapake/vincent-inventory-api/utils/common"
	"github.com/shopspring/decimal"
)

type InventoryApp interface {
	GetStockProductList(c *fiber.Ctx, request GetStockProductListDTI) ([]product.Products, error)
}

type InventoryAppImpl struct {
	repo *repository.Repository
}

func NewInventoryApp(repo *repository.Repository) InventoryApp {
	return &InventoryAppImpl{
		repo: repo,
	}
}

type GetStockProductListDTI struct {
	ProductCode *string
	ProductName *string
	BrandId     *string
	CategoryId  *string
	MinNetPrice *int32
	MaxNetPrice *int32
	InStock     *bool
	Pagination  common.PaginationRequest
}

type GetStockProductListResponse struct {
	Id                   uuid.UUID
	Name                 string
	HighlightCoverImgUrl string
	BrandId              *uuid.UUID
	CategoryId           *uuid.UUID
	Code                 string
	NetPrice             *decimal.Decimal
	Stock                int32
}

func (a *InventoryAppImpl) GetStockProductList(c *fiber.Ctx, request GetStockProductListDTI) ([]product.Products, error) {
	result, err := a.repo.Product.GetProductList(c, repository.GetProductListFilter{
		ProductCode: request.ProductCode,
		ProductName: request.ProductName,
		BrandId:     request.BrandId,
		CategoryId:  request.CategoryId,
		MinNetPrice: request.MinNetPrice,
		MaxNetPrice: request.MaxNetPrice,
		InStock:     request.InStock,
		Pagination:  request.Pagination,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
