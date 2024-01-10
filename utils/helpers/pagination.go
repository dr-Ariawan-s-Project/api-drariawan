package helpers

import "math"

type Pagination struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

func CountPagination(totalRows int64, page int, perPage int) Pagination {
	var totalPages int64 = 1

	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 10
	}

	totalPages = int64(math.Ceil(float64(totalRows) / float64(perPage)))

	if totalPages <= 0 {
		totalPages = 1
	}
	return Pagination{
		Page:         page,
		Limit:        perPage,
		TotalPages:   int(totalPages),
		TotalRecords: totalRows,
	}
}
