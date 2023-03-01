package internal

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
	"strconv"
)

func GetCategories(context *gin.Context) ([]model.Category, helpers.PaginationData) {
	perPage := viper.GetInt("PAGINATION.PER_PAGE")
	page := viper.GetInt("PAGINATION.PAGE")
	sort := viper.GetString("PAGINATION.SORT")
	queries := context.Request.URL.Query()
	pageStr, ok := context.GetQuery("page")

	if ok {
		page, _ = strconv.Atoi(pageStr)
	}

	sortStr, sortOk := context.GetQuery("sort")

	if sortOk {
		sort = sortStr
	}

	perPageStr, perPageOk := context.GetQuery("per_page")

	if perPageOk {
		perPage, _ = strconv.Atoi(perPageStr)
	}

	var categories []model.Category
	var totalItems int64
	searchValue, _ := context.GetQuery("search")
	clauses := helpers.SearchBuilder("name", searchValue)
	queriesMap := helpers.QueryBuilder(queries)
	helpers.DB.Model(&model.Category{}).Count(&totalItems)
	pagination := helpers.GetPaginationData(page, perPage, totalItems, sort, model.Category{})
	helpers.DB.Preload(clause.Associations).Clauses(clauses...).Where(queriesMap).Order("created_at " + sort).Limit(perPage).Offset(pagination.Offset).Find(&categories)
	return categories, pagination
}
