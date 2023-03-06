package helpers

import (
	"net/url"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type ProductConditionObject struct {
	IsProduct bool
	IsAdmin   bool
}

func QueryBuilder(queries url.Values, productCondition ...ProductConditionObject) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range queries {
		if k != "sort" && k != "page" && k != "search" && k != "per_page" {
			if !productCondition[0].IsAdmin && productCondition[0].IsProduct {
				PendingStatus, _ := model.Pending.Value()
				result["status"] = PendingStatus
			}
			result[k] = v
		}
	}
	log.Info("Result: ", result)
	return result
}

func SearchBuilder(fieldName string, value string) []clause.Expression {
	clauses := make([]clause.Expression, 0)
	if fieldName != "" {
		clauses = append(clauses, clause.Like{Column: fieldName, Value: value + "%"})
	}
	return clauses
}
