package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(context *gin.Context) {

	context.JSON(http.StatusCreated, gin.H{"data": ""})
}

func SignOut(context *gin.Context) {

	context.JSON(http.StatusCreated, gin.H{"data": ""})
}
