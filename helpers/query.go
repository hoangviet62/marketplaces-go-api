package helpers

import (
	"net/url"

	"gorm.io/gorm/clause"
)

func QueryBuilder(queries url.Values) map[string]interface{} {
	result := make(map[string]interface{})
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
