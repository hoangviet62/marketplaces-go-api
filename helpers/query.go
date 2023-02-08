package helpers

import (
	"net/url"
)

func QueryBuilder(queries url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range queries {
		if k != "sort" && k != "page" {
			result[k] = v
		}
	}
	return result
}
