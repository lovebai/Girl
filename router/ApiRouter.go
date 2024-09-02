package router

import (
	apiController "Girl/controller/api"

	"github.com/gin-gonic/gin"
)

// 前台路由
func ApiRouter(router *gin.Engine) {
	index := router.Group("/api")
	{
		index.GET("/RandImg", apiController.RandImgUrl)
	}
}
