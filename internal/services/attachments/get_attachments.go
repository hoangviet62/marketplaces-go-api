package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"github.com/spf13/viper"
	"strconv"
)

func GetAttachments(context *gin.Context) ([]model.Attachment, helpers.PaginationData) {
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
	
	var attachments []model.Attachment
	var totalItems int64
	queriesMap := helpers.QueryBuilder(queries)
	helpers.DB.Model(&model.Attachment{}).Count(&totalItems)
	pagination := helpers.GetPaginationData(page, perPage, totalItems, sort, model.Attachment{})
	helpers.DB.Where(queriesMap).Order("created_at " + sort).Limit(perPage).Offset(pagination.Offset).Find(&attachments)
	return attachments, pagination
}
