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
