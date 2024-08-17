package indexController

import (
	"Girl/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// dao.Mgr.GetSetting(&model.Setting{})
	data := dao.Mgr.GetSetting()

	c.HTML(http.StatusOK, "index/index", gin.H{
		"version":  "1.0.0",
		"title":    "首页",
		"siteName": data[1],
	})
}

func Little(c *gin.Context) {
	c.HTML(http.StatusOK, "index/little", gin.H{
		"title": "点点滴滴",
	})

}

func Leaving(c *gin.Context) {
	c.HTML(http.StatusOK, "index/leaving", gin.H{
		"title": "留言",
	})

}

func Photo(c *gin.Context) {
	c.HTML(http.StatusOK, "index/photo", gin.H{
		"title": "相册",
	})

}

func TodoList(c *gin.Context) {
	c.HTML(http.StatusOK, "index/todolist", gin.H{
		"title": "恋爱列表",
	})

}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "index/about", gin.H{
		"title": "关于",
	})

}

func LittlePost(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "index/little-post", gin.H{
		"id":    id,
		"title": "文章页面",
	})
}
