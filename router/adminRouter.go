package router

import (
	adminController "Girl/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	admin := router.Group("/Admin")
	{
		admin.GET("/login", adminController.Login)
		admin.GET("/", adminController.IndexPage)
		admin.GET("/setting", adminController.SettingPage)
		admin.GET("/leaving", adminController.LeavingPage)
		admin.GET("/little", adminController.LittlePage)
		admin.GET("/little/add", adminController.LittleAddPage)
		admin.GET("/photo", adminController.PhotoPage)
		admin.GET("/photo/add", adminController.PhotoAddPage)
		admin.GET("/todolist", adminController.TodoListPage)
		admin.GET("/todolist/add", adminController.TodoListAddPage)
		admin.GET("/about", adminController.AboutPage)
		admin.GET("/userinfo", adminController.UserInfoPage)
		//
		admin.POST("/leaving/del", adminController.DeleteByid)
		admin.POST("/little/del", adminController.DeleteByid)
		admin.POST("/photo/del", adminController.DeleteByid)
		admin.POST("/todolist/del", adminController.DeleteByid)
		//
		admin.POST("/little/add", adminController.AddLittle)
		admin.POST("/photo/add", adminController.AddPhoto)
		admin.POST("/todolist/add", adminController.AddTodoList)

		//
		admin.GET("/little/update/:id", adminController.UpdateLittlePage)
		admin.POST("/little/update", adminController.UpdateLittle)
		admin.GET("/photo/update/:id", adminController.UpdatePhotoPage)
		admin.POST("/photo/update", adminController.UpdateLPhoto)
	}
}
