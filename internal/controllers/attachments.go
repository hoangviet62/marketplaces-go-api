package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
)

func GetAttachments(context *gin.Context) {
	var attachments, pagination = attachmentService.GetAttachments(context)
	context.JSON(http.StatusOK, gin.H{"data": attachments, "pagination": pagination})
}
