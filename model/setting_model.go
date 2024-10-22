package model

type SettingA struct {
	SiteName    string `json:"sitename"`
	SiteLogo    string `json:"sitelogo"`
	SiteDesc    string `json:"sitedesc"`
	Keywords    string `json:"keywords"`
	Discription string `json:"discription"`
}

type SettingB struct {
	BoyName    string `json:"boyname"`
	GirlName   string `json:"girlname"`
	BoyImgUrl  string `json:"boyimgurl"`
	GirlImgUrl string `json:"girlimgurl"`
	StartTime  string `json:"starttime"`
}

type SettingC struct {
	CardOne       string `json:"cardone"`
	CardOneDesc   string `json:"cardonedesc"`
	CardTwo       string `json:"cardtwo"`
	CardTwoDesc   string `json:"cardtwodesc"`
	CardThree     string `json:"cardthree"`
	CardThreeDesc string `json:"cardthreedesc"`
	SiteIcp       string `json:"siteicp"`
	SiteGaIcp     string `json:"sitegaicp"`
	SiteCopyright string `json:"sitecopyright"`
}

type SettingD struct {
	AvaterShowType int    `json:"avatershowtype"`
	LinkType       int    `json:"linktype"`
	LinkApiUrl     string `json:"linkapiurl"`
	LinkApiToken   string `json:"linkapitoken"`
	SiteBlur       int    `json:"siteblur"`
	SitePajx       int    `json:"sitepajx"`
	BgimgUrl       string `json:"bgimgurl"`
	SiteIcon       string `json:"siteicon"`
	Watch          int    `json:"watch"`
}

type SettingE struct {
	LenvingSum int    `json:"lenvingsum"`
	SiteRegexp string `json:"siteregexp"`
	BlockWord  string `json:"blockword"`
}

type SettingF struct {
	SiteHeadStyle string `json:"siteheadstyle"`
	SiteFootJs    string `json:"sitefootjs"`
}

// 内部
type SumCount struct {
	Lenving     int64
	Article     int64
	Photo       int64
	TodoList    int64
	IpBlackList int64
}

var LoggedInUsers = make(map[string]bool)
