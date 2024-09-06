package InstallController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallInfo struct {
	AdminUrl string `form:"adminUrl"`
	SqlFile  string `form:"sqlFile"`
	Salt     string `form:"salt"`
	BoyName  string `form:"boyName" binding:"required"`
	BoyPass  string `form:"boyPass" binding:"required"`
	BoyQq    string `form:"boyQq" binding:"required"`
	GirlName string `form:"girlName" binding:"required"`
	GirlPass string `form:"girlPass" binding:"required"`
	GirlQq   string `form:"girlQq" binding:"required"`
}

func InstallPage(c *gin.Context) {
	c.HTML(http.StatusOK, "install", gin.H{"code": 200, "msg": ""})
}

func Install(c *gin.Context) {
	var info InstallInfo

	if c.ShouldBind(&info) != nil {
		c.HTML(http.StatusOK, "install", gin.H{"code": 201, "msg": "参数有误！"})
	}

	// fmt.Printf("info: %v\n", info)
	// fmt.Printf("len(info.BoyName): %v\n", len(info.BoyName))
	if len(info.BoyName) < 1 || len(info.BoyPass) < 1 || len(info.BoyQq) < 1 || len(info.GirlName) < 1 || len(info.GirlPass) < 1 || len(info.GirlQq) < 1 {
		c.HTML(http.StatusOK, "install", gin.H{"code": 202, "msg": "用户信息必须填写！"})
	}
}
