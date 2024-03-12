package pagination

import (
	"github.com/peammapake/vincent-inventory-api/utils/common"
	"gorm.io/gorm"
)

func PaginationAndSort(db *gorm.DB, paginationRequest common.PaginationRequest) *gorm.DB {
	query := db
	// Apply pagination
	offset := (paginationRequest.Page - 1) * paginationRequest.PageSize
	query = query.Offset(int(offset)).Limit(int(paginationRequest.PageSize))

	// Apply sorting
	for i := range paginationRequest.SortBy {
		sort := paginationRequest.SortBy[i]

		if sort.OrderBy == common.DESC {
			query = query.Order(sort.SortBy + " DESC")
		} else {
			query = query.Order(sort.SortBy + " ASC")
		}
	}

	return query
}
