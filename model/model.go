package model

import "gorm.io/gorm"

//站点设置
type Siteinfo struct {
	OptionId      int
	SiteName      string
	SiteLogo      string
	SiteDesc      string
	SiteBlur      int
	SitePajx      int
	BoyName       string
	GirlName      string
	BoyQq         string
	GirlQq        string
	StartTime     string
	BgimgUrl      string
	CardOne       string
	CardOneDesc   string
	CardTwo       string
	CardTwoDesc   string
	CardThree     string
	CardThreeDesc string
	SiteIcp       string
	SiteGaIcp     string
	SiteCopyright string
	LenvingSum    int
}

//文章
type Article struct {
	ArticleId      int
	ArticleTitle   string
	ArticleContext string
	ArticleAuthor  string
	ArticleTime    string
}

type Lenving struct {
	LenvingId   int
	LenvingName string
	LenvingQq   string
	LenvingText string
	LenvingTime int64
	LenvingIp   string
	LenvingCity string
}

type Photo struct {
	ImgId   int
	ImgText string
	ImgUrl  string
	ImgTime int64
}

func (Siteinfo) TableName() string {
	return "siteinfo"
}

func (Article) TableName() string {
	return "article"
}

func (Lenving) TableName() string {
	return "lenving"
}

func (Photo) TableName() string {
	return "photo"
}

// 后台用户
type User struct {
	gorm.Model

	Username string
	Password string
}
