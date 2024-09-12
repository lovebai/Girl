package router

import (
	"Girl/src"
	"Girl/utlis"
	"html/template"
	"io"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// loadTemplate 加载由 go-assets-builder 嵌入的模板
func loadTemplate(funcMap template.FuncMap) (*template.Template, error) {
	t := template.New("").Funcs(funcMap)
	for name, file := range src.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func Start() {
	gin.SetMode(utlis.GetConfBody().AppMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	funcMap := template.FuncMap{
		"formatDate":               utlis.FormatAsDate,
		"FormatAsTimeAgo":          utlis.FormatAsTimeAgo,
		"ConvertTimestampToString": utlis.ConvertTimestampToString,
		"ToHtml":                   utlis.ToHtml,
		"ToHtmlTable":              utlis.ToHtmlTable}
	templates, err := loadTemplate(funcMap)
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(templates)
	router.StaticFS("/static", src.VFS)

	// dev
	// router.SetFuncMap(funcMap)
	// router.LoadHTMLGlob("templates/**/*")
	// router.Static("/static", "./static")

	//session cookie
	store := cookie.NewStore([]byte("Girl"))
	router.Use(sessions.Sessions("BoyAndGirl", store))

	//默认路由 404
	NoRouter(router)

	//前台路由
	IndexRouter(router)

	//后台路由
	AdminRouter(router)

	//Api路由
	ApiRouter(router)

	//安装页面
	InstallRouter(router)

	utlis.Header()
	utlis.Tip(utlis.GetConfBody().AppPort)
	router.Run(":" + utlis.GetConfBody().AppPort)

}
