package common

type PaginationRequest struct {
	Page     int32             `json:"page" validate:"min=1"`
	PageSize int32             `json:"pageSize" validate:"min=1"`
	SortBy   []*SortingRequest `json:"sortBy"`
}
