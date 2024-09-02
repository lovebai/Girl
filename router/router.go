package router

import (
	"Girl/utlis"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.SetFuncMap(template.FuncMap{
		"formatDate":               utlis.FormatAsDate,
		"FormatAsTimeAgo":          utlis.FormatAsTimeAgo,
		"ConvertTimestampToString": utlis.ConvertTimestampToString})
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	//前台路由
	IndexRouter(router)

	//后台路由
	AdminRouter(router)

	//Api路由
	ApiRouter(router)

	router.Run(":8100")

}
