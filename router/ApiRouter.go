package router

import (
	apiController "Girl/controller/api"

	"github.com/gin-gonic/gin"
)

// 前台路由
func ApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/RandImg", apiController.RandImgUrl)
		api.GET("/BingImg", apiController.BingTodayUrl)
		api.GET("/Player", apiController.Player)
	}
}
