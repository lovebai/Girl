package router

import (
	"Girl/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	//前台路由
	index := router.Group("/")
	{
		index.GET("/", controller.Index)
		index.GET("/Little", controller.Little)
		index.GET("/Lenving", controller.Leaving)
		index.GET("/Photo", controller.Photo)
		index.GET("/TodoList", controller.TodoList)
		index.GET("/About", controller.About)
	}

	//后台路由
	admin := router.Group("/Admin")
	{
		admin.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"name": "admin", "site": "obai.cc"}) })
		admin.GET("/list", nil)
	}

	router.Run(":8100")

}
