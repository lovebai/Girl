package indexController

import (
	"Girl/dao"
	"Girl/model"
	"Girl/utlis"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 查
type Body struct {
	SiteInfo model.Siteinfo
	Man      model.User
	Woman    model.User
}

func getIndexInfo() Body {
	body := Body{
		SiteInfo: dao.Inx.GetSetting(),
		Man:      dao.Inx.GetUserinfoBySex(1),
		Woman:    dao.Inx.GetUserinfoBySex(0),
	}
	return body
}

func Index(c *gin.Context) {
	info := getIndexInfo()
	c.HTML(http.StatusOK, "index/index", gin.H{
		"title": info.SiteInfo.SiteName,
		"info":  info,
	})
}

func Little(c *gin.Context) {
	info := getIndexInfo()
	article := dao.Inx.GetArticleList()
	c.HTML(http.StatusOK, "index/little", gin.H{
		"title":   "点点滴滴 - " + info.SiteInfo.SiteName,
		"info":    info,
		"article": article,
	})

}

func Leaving(c *gin.Context) {
	info := getIndexInfo()
	count := dao.Inx.GetLenvingCount()
	lenvingList := dao.Inx.GetLenvingListLimit(info.SiteInfo.LenvingSum)
	c.HTML(http.StatusOK, "index/leaving", gin.H{
		"title":       "留言 - " + info.SiteInfo.SiteName,
		"info":        info,
		"count":       count,
		"lenvingList": lenvingList,
	})

}

// 添加留言
func LeavingAdd(c *gin.Context) {
	count := dao.Inx.GetLenvingCount()
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
	status := dao.Inx.AddLenving(data)
	// fmt.Println(status)
	if status {
		c.JSON(http.StatusOK, gin.H{"code": 200})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 201})

}

func Photo(c *gin.Context) {
	info := getIndexInfo()
	imgList := dao.Inx.GetPhotoList()
	c.HTML(http.StatusOK, "index/photo", gin.H{
		"title": "相册 - " + info.SiteInfo.SiteName,
		"info":  info,
		"list":  imgList,
	})

}

func TodoList(c *gin.Context) {
	info := getIndexInfo()
	todolist := dao.Inx.GetTodoList()
	c.HTML(http.StatusOK, "index/todolist", gin.H{
		"title":    "恋爱事件 - " + info.SiteInfo.SiteName,
		"info":     info,
		"todolist": todolist,
	})

}

func About(c *gin.Context) {
	info := getIndexInfo()
	about := dao.Inx.GetAbout()
	c.HTML(http.StatusOK, "index/about", gin.H{
		"title": "关于 - " + info.SiteInfo.SiteName,
		"info":  info,
		"about": about,
	})

}

func LittlePost(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"tip": "参数有误！"})
		return
	}
	post := dao.Inx.GetArticle(aid)
	info := getIndexInfo()
	c.HTML(http.StatusOK, "index/little-post", gin.H{
		"id":    id,
		"title": post.ArticleTitle,
		"post":  post,
		"info":  info,
	})
}
