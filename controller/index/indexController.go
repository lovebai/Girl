package indexController

import (
	"Girl/dao"
	"Girl/model"
	"Girl/utlis"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func Index(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	c.HTML(http.StatusOK, "index/index", gin.H{
		"title": siteinfo.SiteName,
		"info":  siteinfo,
	})
}

func Little(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	article := dao.Mgr.GetArticleList()
	c.HTML(http.StatusOK, "index/little", gin.H{
		"title":   "点点滴滴 - " + siteinfo.SiteName,
		"info":    siteinfo,
		"article": article,
	})

}

func Leaving(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	count := dao.Mgr.GetLenvingCount()
	lenvingList := dao.Mgr.GetLenvingList()
	c.HTML(http.StatusOK, "index/leaving", gin.H{
		"title":       "留言 - " + siteinfo.SiteName,
		"info":        siteinfo,
		"count":       count,
		"lenvingList": lenvingList,
	})

}

// 添加留言
func LeavingAdd(c *gin.Context) {
	count := dao.Mgr.GetLenvingCount()
	name := c.PostForm("name")
	qq := c.PostForm("qq")
	text := c.PostForm("text")
	data := model.Lenving{
		LenvingId:   int(count) + 1,
		LenvingName: name,
		LenvingQq:   qq,
		LenvingText: text,
		LenvingTime: time.Now().Unix(),
		LenvingIp:   utlis.GetIPFromRequest(c),
		LenvingCity: utlis.Get_ip_city(utlis.GetIPFromRequest(c)),
	}
	status := dao.Mgr.AddLenving(data)
	// fmt.Println(status)
	if status {
		c.JSON(http.StatusOK, gin.H{"code": 200})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 201})

}

func Photo(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	imgList := dao.Mgr.GetPhotoList()
	c.HTML(http.StatusOK, "index/photo", gin.H{
		"title": "相册 - " + siteinfo.SiteName,
		"info":  siteinfo,
		"list":  imgList,
	})

}

func TodoList(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	todolist := dao.Mgr.GetTodoList()
	c.HTML(http.StatusOK, "index/todolist", gin.H{
		"title":    "恋爱事件 - " + siteinfo.SiteName,
		"info":     siteinfo,
		"todolist": todolist,
	})

}

func About(c *gin.Context) {
	siteinfo := dao.Mgr.GetSetting()
	c.HTML(http.StatusOK, "index/about", gin.H{
		"title": "关于 - " + siteinfo.SiteName,
		"info":  siteinfo,
	})

}

func LittlePost(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"tip": "参数有误！"})
		return
	}
	post := dao.Mgr.GetArticle(aid)
	content := blackfriday.Run([]byte(post.ArticleContext))
	siteinfo := dao.Mgr.GetSetting()
	c.HTML(http.StatusOK, "index/little-post", gin.H{
		"id":      id,
		"title":   post.ArticleTitle,
		"post":    post,
		"centext": template.HTML(content),
		"info":    siteinfo,
	})
}
