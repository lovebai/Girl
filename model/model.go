package model

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
	StartTime     int64
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
	SiteHeadStyle string
	SiteRegexp    string
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

type TodoList struct {
	ListId     int
	ListStatus int
	ListText   string
	ListImgurl string
}

type About struct {
	Id       int
	Title    string
	Aboutimg string
	Info1    string
	Info2    string
	Info3    string
	Btn1     string
	Btn2     string
	Btnx2    string
	Btnf3    string
	Infox1   string
	Infox2   string
	Infox3   string
	Infox4   string
	Infox5   string
	Infox6   string
	Infof1   string
	Infof2   string
	Infof3   string
	Infof4   string
	Infod1   string
	Infod2   string
	Infod3   string
	Infod4   string
	Infod5   string
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

func (TodoList) TableName() string {
	return "todolist"
}

func (About) TableName() string {
	return "about"
}

// 内部
type SumCount struct {
	Lenving  int64
	Article  int64
	Photo    int64
	TodoList int64
}
