package middleware

import (
	"Girl/dao"
	"Girl/utlis"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 登录状态
func LoginStatus(c *gin.Context) {
	// fmt.Printf("utlis.MD5Encrypt(\"123456\"): %v\n", utlis.MD5Encrypt("123456"))
	session := sessions.Default(c)
	jwtToken := session.Get("jwtToken")
	if jwtToken == nil {
		c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		fmt.Printf("没有cookie")
		c.Abort()
		return
	}
	fmt.Printf("jwtToken: %v\n", jwtToken)
	claims, err := utlis.ParseToken(jwtToken.(string))
	if err != nil {
		c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		return
	}
	fmt.Printf("claims: %v\n", claims)

	res, user := dao.Mgr.GetUserinfoByName(claims.Username)
	fmt.Printf("user.Password: %v c %v %v\n", user.Password, claims.Password, res)
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
