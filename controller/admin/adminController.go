package adminController

import (
	"Girl/dao"
	"Girl/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login", gin.H{
		"title": "管理员登录页面",
	})
}

// 获取所有总数相关数据
func getAllCount() model.SumCount {
	lvCount := dao.Mgr.GetLenvingCountSum()
	atCount := dao.Mgr.GetArticleCountSum()
	tlCount := dao.Mgr.GetTodoListCountSum()
	phCount := dao.Mgr.GetPhotoCountSum()
	sum := &model.SumCount{
		Lenving:  lvCount,
		Article:  atCount,
		Photo:    tlCount,
		TodoList: phCount,
	}
	return *sum
}

// 后台管理地址
func getAdminUrl() string {
	return "Admin"
}

func Index(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	leavingList := dao.Mgr.GetLenvingListAdmin()
	c.HTML(http.StatusOK, "admin/index", gin.H{
		"title":       "首页",
		"admin_url":   getAdminUrl(),
		"countSum":    getAllCount(),
		"leavingList": leavingList,
		"info":        siteinfo,
	})
}

func Setting(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/setting", gin.H{
		"title":     "设置",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func Leaving(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	leavingList := dao.Mgr.GetLenvingListAdmin()
	c.HTML(http.StatusOK, "admin/leaving", gin.H{
		"title":       "留言",
		"admin_url":   getAdminUrl(),
		"countSum":    getAllCount(),
		"leavingList": leavingList,
		"info":        siteinfo,
	})
}

func Little(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetArticleListAdmin()
	c.HTML(http.StatusOK, "admin/little", gin.H{
		"title":      "点滴",
		"admin_url":  getAdminUrl(),
		"countSum":   getAllCount(),
		"littleList": list,
		"info":       siteinfo,
	})
}

func LittleAdd(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title":     "新增点滴",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func Photo(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetPhotoListAdmin()
	c.HTML(http.StatusOK, "admin/photo", gin.H{
		"title":     "相册",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"photoList": list,
		"info":      siteinfo,
	})
}

func PhotoAdd(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title":     "添加相册",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func TodoList(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetTodoListAdmin()
	c.HTML(http.StatusOK, "admin/todolist", gin.H{
		"title":     "恋爱清单",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"todoList":  list,
		"info":      siteinfo,
	})
}

func TodoListAdd(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title":     "添加事件",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func About(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	about := dao.Mgr.GetAboutAdmin()
	c.HTML(http.StatusOK, "admin/about", gin.H{
		"title":     "关于",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"about":     about,
		"info":      siteinfo,
	})
}

func UserInfo(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/userInfo", gin.H{
		"title":     "用户信息",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

// 删除
func DeleteByid(c *gin.Context) {
	id := c.PostForm("id")
	text := c.PostForm("text")
	ty := c.PostForm("type")
	if len(text) < 2 || len(ty) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	lid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	var status int64
	switch string(ty) {
	case "leaving":
		status = dao.Mgr.DeleteLenving(lid)
	case "article":
		status = dao.Mgr.DeleteLittle(lid)
	case "photo":
		status = dao.Mgr.DeletePhoto(lid)
	case "todolist":
		status = dao.Mgr.DeleteTodoList(lid)
	default:
	}

	if status == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "删除失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功!"})
	}
}
