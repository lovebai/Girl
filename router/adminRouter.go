package router

import (
	adminController "Girl/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	admin := router.Group("/Admin")
	{
		// admin.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"name": "admin", "site": "obai.cc"}) })
		admin.GET("/login", adminController.Login)
		admin.GET("/", adminController.Index)
		admin.GET("/setting", adminController.Setting)
		admin.GET("/leaving", adminController.Leaving)
		admin.GET("/little", adminController.Little)
		admin.GET("/little/add", adminController.LittleAdd)
		admin.GET("/photo", adminController.Photo)
		admin.GET("/photo/add", adminController.PhotoAdd)
		admin.GET("/todolist", adminController.TodoList)
		admin.GET("/todolist/add", adminController.TodoListAdd)
		admin.GET("/about", adminController.About)
		admin.GET("/userinfo", adminController.UserInfo)
	}
}
