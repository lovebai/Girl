package router

import (
	InstallController "Girl/controller/install"

	"github.com/gin-gonic/gin"
)

// 安装路由
func InstallRouter(router *gin.Engine) {
	router.GET("/install", InstallController.InstallPage)
	router.POST("/install", InstallController.Install)
}
