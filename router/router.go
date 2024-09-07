package router

import (
	"Girl/utlis"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(utlis.GetConfBody().AppMOde)
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.SetFuncMap(template.FuncMap{
		"formatDate":               utlis.FormatAsDate,
		"FormatAsTimeAgo":          utlis.FormatAsTimeAgo,
		"ConvertTimestampToString": utlis.ConvertTimestampToString,
		"ToHtml":                   utlis.ToHtml,
		"ToHtmlTable":              utlis.ToHtmlTable})
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	//session cookie
	store := cookie.NewStore([]byte("Girl"))
	router.Use(sessions.Sessions("BoyAndGirl", store))

	//前台路由
	IndexRouter(router)

	//后台路由
	AdminRouter(router)

	//Api路由
	ApiRouter(router)

	//安装页面
	InstallRouter(router)

	router.Run(":8100")

}
