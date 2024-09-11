package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 404 页面
func NoPage(c *gin.Context) {
	htmlContext := `<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>Like Girl For Golang / <a href="/" style="text-decoration: none;color: #03A9F4;">Go To Home</a></center>
</body>
</html>`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContext))
}

// 无路由 默认路由
func NoRouter(router *gin.Engine) {
	router.NoRoute(NoPage)
}
