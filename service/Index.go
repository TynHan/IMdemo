package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Index
//	@Tags		首页
//	@Success	200	{string}	welcome
//	@Router		/index [get]
func GetIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to my site",
	})
}
