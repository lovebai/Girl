package router

import (
	indexController "Girl/controller/index"

	"github.com/gin-gonic/gin"
)

// 前台路由
func IndexRouter(router *gin.Engine) {
	index := router.Group("/")
	{
		index.GET("/", indexController.Index)
		index.GET("/Little", indexController.Little)
		index.GET("/Little/:id", indexController.LittlePost)
		index.GET("/Lenving", indexController.Leaving)
		index.POST("/LenvingAdd", indexController.LeavingAdd)
		index.GET("/Photo", indexController.Photo)
		index.GET("/TodoList", indexController.TodoList)
		index.GET("/About", indexController.About)
	}
}
