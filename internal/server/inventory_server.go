package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/internal/app"
	exception "github.com/peammapake/vincent-inventory-api/pkg/exeption"
	"github.com/peammapake/vincent-inventory-api/pkg/validators"
	"github.com/peammapake/vincent-inventory-api/utils/common"
	"github.com/shopspring/decimal"
)

type GetStockProductListRequest struct {
	ProductCode *string                  `json:"productCode"`
	ProductName *string                  `json:"productName"`
	BrandId     *string                  `json:"brandId" validate:"omitnil,uuid4"`
	CategoryId  *string                  `json:"categoryId" validate:"omitnil,uuid4"`
	MinNetPrice *int32                   `json:"minNetPrice"`
	MaxNetPrice *int32                   `json:"maxNetPrice"`
	InStock     *bool                    `json:"inStock"`
	Pagination  common.PaginationRequest `json:"pagination"`
}

type SingleProductResponse struct {
	Id                   *uuid.UUID       `json:"id"`
	ProductName          string           `json:"productName"`
	ProductCode          string           `json:"productCode"`
	HighlightCoverImgUrl *string          `json:"highlightCoverImgUrl"`
	BrandId              *uuid.UUID       `json:"brandId"`
	CategoryId           *uuid.UUID       `json:"categoryId"`
	NetPrice             *decimal.Decimal `json:"netPrice"`
	Stock                int32            `json:"stock"`
}

type GetStockProductListResponse struct {
	Products     []*SingleProductResponse  `json:"products"`
	ProductCount int32                     `json:"productCount"`
	Pagination   common.PaginationResponse `json:"pagination"`
}

func (s *Server) GetStockProductList(c *fiber.Ctx) error {
	// Parse JSON request from context
	var request GetStockProductListRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exception.NewValidationError(err))
	}

	// Validate request field validation
	if err := validators.ValidateStruct(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exception.NewValidationError(err))
	}

	fmt.Println(request.BrandId)

	dti := app.GetStockProductListDTI{
		ProductCode: request.ProductCode,
		ProductName: request.ProductName,
		BrandId:     request.BrandId,
		CategoryId:  request.CategoryId,
		MinNetPrice: request.MinNetPrice,
		MaxNetPrice: request.MaxNetPrice,
		InStock:     request.InStock,
		Pagination:  request.Pagination,
	}

	result, exception := s.App.Inventory.GetStockProductList(c, dti)
	if exception != nil {
		return c.Status(http.StatusInternalServerError).JSON(exception)
	}

	itemList := make([]*SingleProductResponse, 0)
	itemCount := int32(0)

	for i := range result {
		item := result[i]

		itemList = append(itemList, &SingleProductResponse{
			Id:                   item.Id,
			ProductName:          item.Name,
			ProductCode:          item.Code,
			HighlightCoverImgUrl: item.HighlightCoverImgUrl,
			BrandId:              item.BrandId,
			CategoryId:           item.CategoryId,
			NetPrice:             item.NetPrice,
			Stock:                item.Stock,
		})
		itemCount++
	}

	response := GetStockProductListResponse{
		Products:     itemList,
		ProductCount: itemCount,
		Pagination:   common.PaginationResponse{},
	}

	return c.JSON(response)
}
