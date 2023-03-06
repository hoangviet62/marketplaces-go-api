package helpers

import (
	"net/url"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

type ProductConditionObject struct {
	IsProduct bool
	IsAdmin   bool
}

func QueryBuilder(queries url.Values, productCondition ...ProductConditionObject) map[string]interface{} {
	result := make(map[string]interface{})

	if productCondition != nil && !productCondition[0].IsAdmin && productCondition[0].IsProduct {
		ApprovedStatus, _ := model.Approved.Value()
		result["status"] = ApprovedStatus
	}

	for k, v := range queries {
		if k != "sort" && k != "page" && k != "search" && k != "per_page" {
			result[k] = v
		}
	}
	return result
}

func SearchBuilder(fieldName string, value string) []clause.Expression {
	clauses := make([]clause.Expression, 0)
	if fieldName != "" {
		clauses = append(clauses, clause.Like{Column: fieldName, Value: value + "%"})
	}
	return clauses
}
