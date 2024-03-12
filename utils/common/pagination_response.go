package common

type PaginationResponse struct {
	TotalRecords int32 `json:"totalRecords"`
	TotalPages   int32 `json:"totalPages"`
	Page         int32 `json:"page"`
	PageSize     int32 `json:"pageSize"`
}
