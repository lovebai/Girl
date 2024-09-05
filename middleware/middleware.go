package middleware

import (
	"Girl/dao"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginStatus(c *gin.Context) {
	// fmt.Printf("utlis.MD5Encrypt(\"123456\"): %v\n", utlis.MD5Encrypt("123456"))
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		fmt.Printf(" !username: %v\n", username)
		c.Redirect(http.StatusMovedPermanently, "/Admin/login")
		return
	}
	res, _ := dao.Mgr.GetUserinfoByName(username.(string))

	if res == 0 {
		c.Redirect(http.StatusMovedPermanently, "/Admin/login")
		return
	}

	// fmt.Printf("username: %v\n", username)
}
