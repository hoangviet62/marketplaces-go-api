package helpers

import (
	"gorm.io/gorm/clause"
	"net/url"
)

func QueryBuilder(queries url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range queries {
		if k != "sort" && k != "page" && k != "search" && k != "perPage" {
			result[k] = v
		}
	}
	return result
}

func SearchBuilder(fieldName string, value string) []clause.Expression {
	clauses := make([]clause.Expression, 0)
	clauses = append(clauses, clause.Like{Column: fieldName, Value: value + "%"})
	return clauses
}
