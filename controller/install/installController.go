package InstallController

import (
	"Girl/dao"
	"Girl/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallInfo struct {
	BoyName  string `form:"boyName" binding:"required"`
	BoyPass  string `form:"boyPass" binding:"required"`
	BoyQq    string `form:"boyQq" binding:"required"`
	GirlName string `form:"girlName" binding:"required"`
	GirlPass string `form:"girlPass" binding:"required"`
	GirlQq   string `form:"girlQq" binding:"required"`
}

func InstallPage(c *gin.Context) {
	if dao.Ist.GetWebSiteInfo() != 0 {
		c.JSON(http.StatusOK, gin.H{"code": 205, "msg": "抱歉，您已安装成功，直接访问即可。如需重新安装请删除数据库文件"})
	} else {
		c.HTML(http.StatusOK, "install", gin.H{"code": 200, "msg": ""})
	}
}

func Install(c *gin.Context) {
	var info InstallInfo

	if c.ShouldBind(&info) != nil {
		c.HTML(http.StatusOK, "install", gin.H{"code": 201, "msg": "参数有误！"})
		return
	}

	if len(info.BoyName) < 1 || len(info.BoyPass) < 1 || len(info.BoyQq) < 1 || len(info.GirlName) < 1 || len(info.GirlPass) < 1 || len(info.GirlQq) < 1 {
		c.HTML(http.StatusOK, "install", gin.H{"code": 202, "msg": "用户信息必须填写！"})
		return
	}

	//数据库
	if dao.Ist.CreateSiteinfo() == 0 || dao.Ist.CreateAbout() == 0 || dao.Ist.CreateArticle(info.BoyName) == 0 || dao.Ist.CreateLenving() == 0 {
		c.HTML(http.StatusOK, "install", gin.H{"code": 206, "msg": "抱歉安装失败！"})
		return
	} else if dao.Ist.CreatePhoto() == 0 || dao.Ist.CreateIpBlackList() == 0 || dao.Ist.CreateTodoList() == 0 {
		c.HTML(http.StatusOK, "install", gin.H{"code": 206, "msg": "抱歉安装失败！"})
		return
	} else if dao.Ist.CreateUser(1, info.BoyName, utlis.MD5Encrypt(info.BoyPass), info.BoyQq, 1) == 0 || dao.Ist.CreateUser(2, info.GirlName, utlis.MD5Encrypt(info.GirlPass), info.GirlQq, 0) == 0 {
		c.HTML(http.StatusOK, "install", gin.H{"code": 206, "msg": "抱歉安装失败！"})
		return
	} else {
		// c.HTML(http.StatusOK, "install", gin.H{"code": 200, "msg": "恭喜您安装成功啦！"})
		c.Redirect(http.StatusMovedPermanently, "/")
	}

}
