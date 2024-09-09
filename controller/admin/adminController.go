package adminController

import (
	"Girl/dao"
	"Girl/model"
	"Girl/utlis"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取所有总数相关数据
func getAllCount() model.SumCount {
	lvCount := dao.Mgr.GetLenvingCountSum()
	atCount := dao.Mgr.GetArticleCountSum()
	tlCount := dao.Mgr.GetTodoListCountSum()
	phCount := dao.Mgr.GetPhotoCountSum()
	blCount := dao.Mgr.GetIpBlackListCountSum()
	sum := &model.SumCount{
		Lenving:     lvCount,
		Article:     atCount,
		Photo:       phCount,
		TodoList:    tlCount,
		IpBlackList: blCount,
	}
	return *sum
}

// 后台管理地址
func getAdminUrl() string {
	return utlis.GetConfBody().Path
}

// 根据jwt获取用户
func getUserinfoByJwt(c *gin.Context) model.User {
	session := sessions.Default(c)
	jwtToken := session.Get("jwtToken")
	user := model.User{}
	if jwtToken == nil {
		// c.Redirect(http.StatusFound, "/"+utlis.GetConfBody().Path+"/login")
		return model.User{}
	}
	claims, err := utlis.ParseToken(jwtToken.(string))
	if err != nil {
		return model.User{}
	}

	user.Username = claims.Username
	user.Password = claims.Password
	return user
}

// 后台用户
func getuserinfo(c *gin.Context) model.User {
	username := getUserinfoByJwt(c).Username
	u := model.User{}
	if len(username) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/"+getAdminUrl()+"/login")
		return u
	}
	res, user := dao.Mgr.GetUserinfoByName(username)
	if res == 0 {
		return u
	}
	return user
}

func LoginPage(c *gin.Context) {
	un := getUserinfoByJwt(c).Username
	if len(un) == 0 {
		res, _ := dao.Mgr.GetUserinfoByName(un)
		if res != 0 {
			c.Redirect(http.StatusMovedPermanently, "")
			return
		}
	}
	siteinfo := dao.Mgr.GetSettingInfo()
	c.HTML(http.StatusOK, "admin/login", gin.H{
		"title":     "管理员登录页面",
		"admin_url": getAdminUrl(),
		"info":      siteinfo,
	})
}

func IndexPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	leavingList := dao.Mgr.GetLenvingListAdminLimit(10)
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/index", gin.H{
		"title":       "首页",
		"admin_url":   getAdminUrl(),
		"countSum":    getAllCount(),
		"leavingList": leavingList,
		"info":        siteinfo,
		"user":        user,
	})
}

func SettingPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/setting", gin.H{
		"title":     "设置",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
	})
}

func LeavingPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	leavingList := dao.Mgr.GetLenvingListAdmin()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/leaving", gin.H{
		"title":       "留言",
		"admin_url":   getAdminUrl(),
		"countSum":    getAllCount(),
		"leavingList": leavingList,
		"info":        siteinfo,
		"user":        user,
	})
}

func LittlePage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetArticleListAdmin()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/little", gin.H{
		"title":      "点滴",
		"admin_url":  getAdminUrl(),
		"countSum":   getAllCount(),
		"littleList": list,
		"info":       siteinfo,
		"user":       user,
	})
}

func LittleAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	userAll := dao.Mgr.GetAllUserList()
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title":     "新增点滴",
		"h4":        "新增",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
		"userlist":  userAll,
	})
}

func PhotoPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetPhotoListAdmin()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/photo", gin.H{
		"title":     "相册",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"photoList": list,
		"info":      siteinfo,
		"user":      user,
	})
}

func PhotoAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title":     "添加相册",
		"h4":        "新增",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
		"ph":        model.Photo{ImgTime: utlis.GetTimeUnix()},
	})
}

func TodoListPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	list := dao.Mgr.GetTodoListAdmin()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/todolist", gin.H{
		"title":     "恋爱清单",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"todoList":  list,
		"info":      siteinfo,
		"user":      user,
	})
}

func IpBlackListPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	list := dao.Mgr.GetIpBlackListAdmin()
	c.HTML(http.StatusOK, "admin/blacklist", gin.H{
		"title":     "IP黑名单列表",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
		"iplist":    list,
	})
}

func IpBlackListAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/blacklistadd", gin.H{
		"title":     "IP黑名单列表",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
	})
}

func TodoListAddPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title":     "添加事件",
		"h4":        "添加",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
	})
}

func AboutPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	about := dao.Mgr.GetAboutAdmin()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/about", gin.H{
		"title":     "关于",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"about":     about,
		"info":      siteinfo,
		"user":      user,
	})
}

func UserInfoPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/userInfo", gin.H{
		"title":     "用户信息",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
	})
}

func OtherSetPage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/otherSet", gin.H{
		"title":     "其他设置",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"info":      siteinfo,
		"user":      user,
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
	user := getuserinfo(c)
	userAll := dao.Mgr.GetAllUserList()
	c.HTML(http.StatusOK, "admin/littleAdd", gin.H{
		"title":     "修改点滴",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"art":       art,
		"info":      siteinfo,
		"user":      user,
		"userlist":  userAll,
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
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/photoAdd", gin.H{
		"title":     "修改图片",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"ph":        ph,
		"info":      siteinfo,
		"user":      user,
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
	user := getuserinfo(c)
	c.HTML(http.StatusOK, "admin/todoListAdd", gin.H{
		"title":     "添加事件",
		"h4":        "修改",
		"admin_url": getAdminUrl(),
		"countSum":  getAllCount(),
		"tl":        tl,
		"info":      siteinfo,
		"user":      user,
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
	case "blacklist":
		status = dao.Mgr.DeleteIpBlackList(lid)
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

// 增加新IP
func AddIpBlackList(c *gin.Context) {
	ip := c.PostForm("ip")
	text := c.PostForm("text")
	if len(text) < 1 || len(ip) < 2 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
		return
	}

	id := getAllCount().IpBlackList + 1
	if dao.Mgr.AddIpBlackList(int(id), ip, text) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "新增IP黑名单失败!"})
	} else {
		// c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新增成功!"})
		c.Redirect(http.StatusMovedPermanently, "/"+getAdminUrl()+"/blacklist")
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

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "用户名和密码不能为空！"})
		return
	}

	state, user := dao.Mgr.GetUserinfoByName(username)
	if state == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "用户名不存在！"})
		return
	}

	if user.Password != utlis.MD5Encrypt(password) {
		c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "密码错误！"})
		return
	}

	session := sessions.Default(c)
	session.Options(sessions.Options{Path: "/", MaxAge: 3600 * 24}) //12小时过期
	// 生成 JWT
	usernameJwt, err := utlis.GenerateJWT(user.Username, user.Password, 12*60) //12小时过期
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 210, "msg": "could not create token"})
		return
	}
	session.Set("jwtToken", usernameJwt)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "恭喜，登录成功啦！"})
}

// 退出登录
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: "/", MaxAge: -1}) //清除
	session.Set("jwtToken", 0)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/"+getAdminUrl()+"/login")
}

// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	uid := c.PostForm("id")
	uname := c.PostForm("username")
	upwd := c.PostForm("password")
	cupwd := c.PostForm("cpassword")
	uqq := c.PostForm("qq")
	id, err1 := strconv.Atoi(uid)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "传递参数有误！"})
		return
	}

	session := sessions.Default(c)
	username := getUserinfoByJwt(c).Username
	oldpwd := getUserinfoByJwt(c).Password
	// _, user := dao.Mgr.GetUserinfoByName(username.(string))
	if upwd != cupwd {
		c.JSON(http.StatusOK, gin.H{"code": 203, "msg": "两个密码不一样！"})
		return
	}

	if username != uname {
		if len(uname) < 1 {
			c.JSON(http.StatusOK, gin.H{"code": 204, "msg": "传递参数有误！"})
			return
		} else {
			//改用户名
			if dao.Mgr.UpdateUsername(id, uname, uqq) == 0 {
				c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "修改失败!"})
				return
			} else {
				if oldpwd != utlis.MD5Encrypt(cupwd) && len(cupwd) != 0 {
					//改
					if len(cupwd) < 5 {
						c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "密码不能小于6位数！"})
						return
					} else {
						if dao.Mgr.UpdateNamePassQQ(id, uname, utlis.MD5Encrypt(cupwd), uqq) == 0 {
							c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "修改失败!"})
							return
						} else {
							c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "用户名和密码修改成功!"})
							session.Options(sessions.Options{Path: "/", MaxAge: -1}) //清除
							// session.Set("password", utlis.MD5Encrypt(cupwd))
							session.Set("jwtToken", 0)
							session.Save()
							return
						}
					}

				} else {
					c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "用户名修改成功!"})
					session.Options(sessions.Options{Path: "/", MaxAge: -1}) //清除
					session.Set("jwtToken", 0)
					session.Save()
					return
				}
			}
		}
	}

	if oldpwd != utlis.MD5Encrypt(cupwd) && len(cupwd) != 0 {
		//改
		if len(cupwd) < 5 {
			c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "密码不能小于6位数！"})
			return
		} else {
			if dao.Mgr.UpdatePassword(id, utlis.MD5Encrypt(cupwd), uqq) == 0 {
				c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "修改失败!"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "密码修改成功!"})
				session.Options(sessions.Options{Path: "/", MaxAge: -1}) //清除
				session.Set("jwtToken", 0)
				session.Save()
				return
			}
		}

	}

	if dao.Mgr.UpdateQQ(id, uqq) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 202, "msg": "修改失败!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "QQ账号修改成功!"})
	}

}

// 信息返回
func resultMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

// 上传图片
func UploadImage(c *gin.Context) {
	siteinfo := dao.Mgr.GetSettingInfo()
	file, err := c.FormFile("file")
	if err != nil {
		resultMsg(c, 205, "参数错误！")
		return
	}
	// 创建一个新的表单数据
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 将文件写入表单数据
	part, err1 := writer.CreateFormFile("image", file.Filename)
	if err1 != nil {
		resultMsg(c, 206, "表单写入有误！")
		return
	}

	// 读取文件内容并写入到表单
	fileContent, err2 := file.Open()
	if err2 != nil {
		resultMsg(c, 205, "参数错误！")
		return
	}

	defer fileContent.Close()

	fileBytes, err3 := io.ReadAll(fileContent)
	if err3 != nil {
		resultMsg(c, 205, "参数错误！")
		return
	}

	part.Write(fileBytes)

	err4 := writer.WriteField("token", siteinfo.LinkApiToken)
	if err4 != nil {
		resultMsg(c, 206, "表单写入有误！")
		return
	}

	writer.Close()

	// 创建一个新的请求，转发到图片服务器
	req, err5 := http.NewRequest("POST", siteinfo.LinkApiUrl, &buf)
	if err5 != nil {
		resultMsg(c, 204, "请求图床异常！")
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求到图片服务器
	client := &http.Client{}
	resp, err6 := client.Do(req)
	if err6 != nil {
		resultMsg(c, 203, "图片上传失败！")
		return
	}
	defer resp.Body.Close()

	// 返回图片服务器的响应
	responseBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)

}
