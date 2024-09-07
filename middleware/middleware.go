package middleware

import (
	"Girl/dao"
	"Girl/utlis"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登录状态
func LoginStatus(c *gin.Context) {
	// fmt.Printf("utlis.MD5Encrypt(\"123456\"): %v\n", utlis.MD5Encrypt("123456"))
	session := sessions.Default(c)
	username := session.Get("username")
	pwd := session.Get("password")
	if username == nil {
		// fmt.Printf(" !username: %v\n", username)
		c.Redirect(http.StatusMovedPermanently, "/"+utlis.GetConfBody().Path+"/login")
		return
	}
	res, user := dao.Mgr.GetUserinfoByName(username.(string))

	if res == 0 {
		c.Redirect(http.StatusMovedPermanently, "/"+utlis.GetConfBody().Path+"/login")
		return
	}

	if user.Password != pwd {
		c.Redirect(http.StatusMovedPermanently, "/"+utlis.GetConfBody().Path+"/login")
	}

	// fmt.Printf("username: %v\n", username)
}

// 跳转安装
func GotoInstall(c *gin.Context) {
	// fmt.Printf("dao.Ist.GetWebSiteInfo(): %v\n", dao.Ist.GetWebSiteInfo())
	if dao.Ist.GetWebSiteInfo() == 0 {
		c.Redirect(http.StatusMovedPermanently, "/install")
		return
	}
}
