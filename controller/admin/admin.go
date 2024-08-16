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
