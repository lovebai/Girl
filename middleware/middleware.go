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
	session := sessions.Default(c)
	jwtToken := session.Get("jwtToken")
	if jwtToken == nil {
		c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		c.Abort()
		return
	}
	claims, err := utlis.ParseToken(jwtToken.(string))
	if err != nil {
		c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		c.Abort()
		return
	}
	res, user := dao.Mgr.GetUserinfoByName(claims.Username)
	if res == 0 || user.Password != claims.Password {
		c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		c.Abort()
		return
	}
	c.Next()
}

// 跳转安装
func GotoInstall(c *gin.Context) {
	if dao.Ist.GetWebSiteInfo() == 0 {
		c.Redirect(http.StatusFound, "/install")
		return
	}
}

// ip地址拦截
func FirewallsMiddleware(c *gin.Context) {
	clientIp := utlis.GetIPFromRequest(c)
	row, _ := dao.Inx.GetBlackByIp(clientIp)
	if row == 1 {
		c.Redirect(http.StatusMovedPermanently, "/firewalls")
		return
	}
}
