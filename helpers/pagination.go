package helpers

import (
	"math"
)

type PaginationData struct {
	NextPage     int    `json:"next_page"`
	PreviousPage int    `json:"previous_page"`
	CurrentPage  int    `json:"current_page"`
	TotalPages   int    `json:"total_pages"`
	Offset       int    `json:"offset"`
	TotalItems   int64  `json:"total_items"`
	Sort         string `json:"sort"`
}

func GetPaginationData(page int, perPage int, totalItems int64, sort string, model interface{}) PaginationData {
	// Calculate Total Pages
	var totalRows int64
	DB.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows) / float64(perPage))

	// Calculate offset
	offset := (page - 1) * perPage

	return PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		Offset:       offset,
		Sort:         sort,
		TotalPages:   int(totalPages),
		TotalItems:   totalItems,
	}
}
