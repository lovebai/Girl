package router

import (
	adminController "Girl/controller/admin"
	"Girl/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	//后台登录
	router.GET("/Admin/login", adminController.LoginPage)
	router.POST("/Admin/login", adminController.Login)

	//登录成功
	admin := router.Group("/Admin", middleware.LoginStatus)
	{

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
		admin.GET("/otherset", adminController.OtherSetPage)
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
		admin.POST("/photo/update", adminController.UpdatePhoto)
		admin.GET("/todolist/update/:id", adminController.UpdateTodoListPage)
		admin.POST("/todolist/update", adminController.UpdateTolist)
		admin.POST("/about/update", adminController.UpdateAbout)
		admin.POST("/siteinfo/updateA", adminController.UpdateSiteInfoA)
		admin.POST("/siteinfo/updateB", adminController.UpdateSiteInfoB)
		admin.POST("/siteinfo/updateC", adminController.UpdateSiteInfoC)
		admin.POST("/siteinfo/updateD", adminController.UpdateSiteInfoD)
		admin.POST("/siteinfo/updateE", adminController.UpdateSiteInfoE)
		admin.POST("/siteinfo/updateF", adminController.UpdateSiteInfoF)
		admin.POST("/userinfo/update", adminController.UpdateUserInfo)

		admin.GET("/logout", adminController.Logout)

		admin.POST("/upload/image", adminController.UploadImage)

	}
}
