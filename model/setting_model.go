package model

type SettingA struct {
	SiteName string `json:"sitename"`
	SiteLogo string `json:"sitelogo"`
	SiteDesc string `json:"sitedesc"`
	SiteBlur int    `json:"siteblur"`
	SitePajx int    `json:"sitepajx"`
}

type SettingB struct {
	BoyName   string `json:"boyname"`
	GirlName  string `json:"girlname"`
	BoyQq     string `json:"boyqq"`
	GirlQq    string `json:"girlqq"`
	StartTime string `json:"starttime"`
}

type SettingC struct {
	BgimgUrl      string `json:"bgimgurl"`
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
}

type SettingE struct {
	LenvingSum int    `json:"lenvingsum"`
	SiteRegexp string `json:"siteregexp"`
	BlockWord  string `json:"blockword"`
}

type SettingF struct {
	SiteHeadStyle string `json:"siteheadstyle"`
	SiteHeadLink  string `json:"siteheadlink"`
	SiteFootJs    string `json:"sitefootjs"`
	SiteFootLink  string `json:"sitefootlink"`
}
