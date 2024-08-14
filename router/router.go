package router

import (
	"Girl/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/static", "./static")

	//前台路由
	e.GET("/", controller.Index)
	e.GET("/Little", controller.Little)
	e.GET("/Lenving", controller.Leaving)
	e.GET("/Photo", controller.Photo)
	e.GET("/TodoList", controller.TodoList)
	e.GET("/About", controller.About)

	e.Run(":8100")

}
