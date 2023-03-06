package internal

import (
	// "fmt"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	userService "github.com/hoangviet62/marketplaces-go-api/internal/services/users"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

func GetProducts(context *gin.Context) ([]model.Product, helpers.PaginationData) {
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

	productCondition := helpers.ProductConditionObject{}
	token := context.GetHeader("Authorization")

	if token == "" {
		productCondition = helpers.ProductConditionObject{IsProduct: true, IsAdmin: false}
	} else {
		user, _ := userService.GetCurrentUser(context)
		userRole, _ := user.Role.Value()
		adminRole, _ := model.Admin.Value()
		productCondition = helpers.ProductConditionObject{IsProduct: true, IsAdmin: userRole == adminRole}
	}

	fmt.Println("IsProduct: ", productCondition.IsProduct)
	fmt.Println("IsAdmin: ", productCondition.IsAdmin)

	products := []model.Product{}

	var totalItems int64
	searchValue, _ := context.GetQuery("search")
	clauses := helpers.SearchBuilder("name", searchValue)
	queriesMap := helpers.QueryBuilder(queries, productCondition)
	helpers.DB.Model(&model.Product{}).Count(&totalItems)
	pagination := helpers.GetPaginationData(page, perPage, totalItems, sort, model.Product{})
	helpers.DB.Preload(clause.Associations).Clauses(clauses...).Where(queriesMap).Order("created_at " + sort).Limit(perPage).Offset(pagination.Offset).Find(&products)
	return products, pagination
}
