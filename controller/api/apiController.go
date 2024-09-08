package apiController

import (
	"Girl/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RandImgUrl(c *gin.Context) {
	arr := []string{
		"http://qzonestyle.gtimg.cn/qzone/qzactStatics/imgs/20171122191532_f2975b.jpg",
		"http://qzonestyle.gtimg.cn/qzone/qzactStatics/imgs/20171123181522_c48800.jpg",
		"http://qzonestyle.gtimg.cn/qzone/qzactStatics/imgs/20171122191603_896cd9.jpg",
		"http://qzonestyle.gtimg.cn/qzone/qzactStatics/imgs/20171122191630_ff8fef.jpg",
		"https://isux.tencent.com/static/images/resources/banner-1.jpg",
		"https://isux.tencent.com/static/images/resources/banner-2.jpg",
		"https://isux.tencent.com/static/images/brands/brands-1.jpg"}
	url := utlis.GetRandomElement(arr)
	c.Redirect(http.StatusMovedPermanently, url)
}

func BingTodayUrl(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://api.bducds.com/api/bing_images/")
}
