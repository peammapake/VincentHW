package common

type SortingRequest struct {
	SortBy  string `json:"sortBy" validate:"required"`
	OrderBy Order  `json:"orderBy" validate:"required,oneof=ASC DESC"`
}

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)
