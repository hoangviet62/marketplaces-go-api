package internal

import (
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

func GetOrderItems(context *gin.Context) ([]model.OrderItem, helpers.PaginationData) {
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

	orderItems := []model.OrderItem{}
	var totalItems int64
	searchValue, _ := context.GetQuery("search")
	clauses := helpers.SearchBuilder("", searchValue)
	queriesMap := helpers.QueryBuilder(queries)
	helpers.DB.Model(&model.OrderItem{}).Count(&totalItems)
	pagination := helpers.GetPaginationData(page, perPage, totalItems, sort, model.OrderItem{})
	helpers.DB.Preload(clause.Associations).Clauses(clauses...).Where(queriesMap).Order("created_at " + sort).Limit(perPage).Offset(pagination.Offset).Find(&orderItems)
	return orderItems, pagination
}
