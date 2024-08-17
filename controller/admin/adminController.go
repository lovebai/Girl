package adminController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login", gin.H{
		"title": "管理员登录页面",
	})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index", gin.H{
		"title": "首页",
	})
}

func Setting(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/setting", gin.H{
		"title": "设置",
	})
}

func Leaving(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/leaving", gin.H{
		"title": "留言",
	})
}

func Little(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/little", gin.H{
		"title": "点滴",
	})
}

func LittleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title": "新增点滴",
	})
}

func Photo(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/photo", gin.H{
		"title": "相册",
	})
}

func PhotoAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title": "添加相册",
	})
}

func TodoList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/todolist", gin.H{
		"title": "清单",
	})
}

func TodoListAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title": "添加事件",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/about", gin.H{
		"title": "关于",
	})
}
