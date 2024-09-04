package adminController

import (
	"Girl/dao"
	"Girl/model"
	"Girl/utlis"
	"fmt"
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
		Photo:    phCount,
		TodoList: tlCount,
	}
	return *sum
}

// 后台管理地址
func getAdminUrl() string {
	return "Admin"
}

func IndexPage(c *gin.Context) {
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

func SettingPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/setting", gin.H{
		"title":     "设置",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func LeavingPage(c *gin.Context) {
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

func LittlePage(c *gin.Context) {
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

func LittleAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title":     "新增点滴",
		"h4":        "新增",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func PhotoPage(c *gin.Context) {
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

func PhotoAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title":     "添加相册",
		"h4":        "新增",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func TodoListPage(c *gin.Context) {
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

func TodoListAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title":     "添加事件",
		"h4":        "添加",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func AboutPage(c *gin.Context) {
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

func UserInfoPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/userInfo", gin.H{
		"title":     "用户信息",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func OtherSetPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/otherSet", gin.H{
		"title":     "其他设置",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
	})
}

func UpdateLittlePage(c *gin.Context) {
	id := c.Param("id")
	lid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	siteinfo := dao.Mgr.GetSettingInfo()
	art := dao.Mgr.GetArticleAdminByID(lid)
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title":     "修改点滴",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"art":       art,
		"info":      siteinfo,
	})
}

func UpdatePhotoPage(c *gin.Context) {
	id := c.Param("id")
	lid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	siteinfo := dao.Mgr.GetSettingInfo()
	ph := dao.Mgr.GetPhotoAdminByID(lid)
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title":     "修改图片",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"ph":        ph,
		"info":      siteinfo,
	})
}

func UpdateTodoListPage(c *gin.Context) {
	id := c.Param("id")
	lid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	siteinfo := dao.Mgr.GetSettingInfo()
	tl := dao.Mgr.GetTodoListAdminByID(lid)
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title":     "添加事件",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"tl":        tl,
		"info":      siteinfo,
	})
}

/**  删除 **/

// 根据ID删除
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
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
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

/** 新增内容 **/

// 增加新文章
func AddLittle(c *gin.Context) {
	articletitle := c.PostForm("articletitle")
	articletext := c.PostForm("articletext")
	articlename := c.PostForm("articlename")
	if len(articlename) < 2 || len(articletext) < 2 || len(articletitle) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	id := getAllCount().Article + 1
	if dao.Mgr.AddLittle(int(id), articlename, articletitle, articletext) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "新增文章失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新增文章成功!"})
	}
}

// 增加新相册
func AddPhoto(c *gin.Context) {
	imgDatd := c.PostForm("imgDatd")
	imgText := c.PostForm("imgText")
	imgUrl := c.PostForm("imgUrl")
	if len(imgDatd) < 2 || len(imgText) < 2 || len(imgUrl) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	id := getAllCount().Photo + 1
	if dao.Mgr.AddPhoto(int(id), imgDatd, imgText, imgUrl) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "新增图片失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新增图片成功!"})
	}
}

// 增加新事件
func AddTodoList(c *gin.Context) {
	text := c.PostForm("text")
	state := c.PostForm("state")
	imgurl := c.PostForm("imgurl")
	if len(text) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	st, err := strconv.Atoi(state)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}
	id := getAllCount().TodoList + 1
	if dao.Mgr.AddTodoList(int(id), st, text, imgurl) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "新增文章失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新增文章成功!"})
	}
}

/**  更新 **/

// Post 更新点滴（文章）
func UpdateLittle(c *gin.Context) {
	articleid := c.PostForm("id")
	articletitle := c.PostForm("articletitle")
	articletext := c.PostForm("articletext")
	if len(articletext) < 2 || len(articletitle) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	aid, err := strconv.Atoi(articleid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateLittles(int(aid), articletitle, articletext) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改文章失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改文章成功!"})
	}
}

// Post 更新图片
func UpdatePhoto(c *gin.Context) {
	pid := c.PostForm("id")
	pt := c.PostForm("imgText")
	pd := c.PostForm("imgDatd")
	pu := c.PostForm("imgUrl")
	fmt.Println("====", pid, pt, pd, pu)
	if len(pt) < 2 || len(pd) < 2 || len(pu) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	aid, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}
	ph := model.Photo{
		ImgId:   aid,
		ImgText: pt,
		ImgUrl:  pu,
		ImgTime: utlis.ConvertToTimestamp(string(pd)),
	}
	if dao.Mgr.UpdatePhotos(ph) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新清单
func UpdateTolist(c *gin.Context) {
	tid := c.PostForm("id")
	text := c.PostForm("text")
	state := c.PostForm("state")
	imgurl := c.PostForm("imgurl")
	if len(text) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	id, err := strconv.Atoi(tid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}
	st, err1 := strconv.Atoi(state)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}
	tl := model.TodoList{
		ListId:     id,
		ListStatus: st,
		ListText:   text,
		ListImgurl: imgurl,
	}
	if dao.Mgr.UpdateTodolists(tl) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新关于对话
func UpdateAbout(c *gin.Context) {
	var data model.About
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateAbouts(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}

}

// Post 更新设置A
func UpdateSiteInfoA(c *gin.Context) {
	var data model.SettingA
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingA(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新设置B
func UpdateSiteInfoB(c *gin.Context) {
	var data model.SettingB
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingB(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新设置C
func UpdateSiteInfoC(c *gin.Context) {
	var data model.SettingC
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingC(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新设置D
func UpdateSiteInfoD(c *gin.Context) {
	var data model.SettingD
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingD(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新设置E
func UpdateSiteInfoE(c *gin.Context) {
	var data model.SettingE
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingE(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}

// Post 更新设置F
func UpdateSiteInfoF(c *gin.Context) {
	var data model.SettingF
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}
	if dao.Mgr.UpdateSettingF(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功!"})
	}
}
