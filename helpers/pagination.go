package helpers

import (
	"math"
)

type PaginationData struct {
	NextPage     int    `json:"nextPage"`
	PreviousPage int    `json:"previousPage"`
	CurrentPage  int    `json:"currentPage"`
	TotalPages   int    `json:"totalPages"`
	Offset       int    `json:"offset"`
	TotalItems   int64  `json:"totalItems"`
	Sort         string `json:"sort"`
}

func GetPaginationData(page int, perPage int, totalItems int64, sort string, model interface{}) PaginationData {
	// Calculate Total Pages
	var totalRows int64
	DB.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

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
