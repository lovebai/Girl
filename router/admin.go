package router

import "github.com/gin-gonic/gin"

func AdminRouter(router *gin.Engine) {
	admin := router.Group("/Admin")
	{
		admin.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"name": "admin", "site": "obai.cc"}) })
		admin.GET("/list", nil)
	}
}
