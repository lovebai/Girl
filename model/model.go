package model

//站点设置
type Siteinfo struct {
	OptionId       int
	SiteName       string
	SiteLogo       string
	SiteDesc       string
	SiteBlur       int
	SitePajx       int
	BoyName        string
	GirlName       string
	BoyImgUrl      string
	GirlImgUrl     string
	StartTime      int64
	BgimgUrl       string
	CardOne        string
	CardOneDesc    string
	CardTwo        string
	CardTwoDesc    string
	CardThree      string
	CardThreeDesc  string
	SiteIcp        string
	SiteGaIcp      string
	SiteCopyright  string
	LenvingSum     int
	SiteHeadStyle  string
	SiteFootJs     string
	SiteRegexp     string
	BlockWord      string
	AvaterShowType int
	LinkType       int
	LinkApiUrl     string
	LinkApiToken   string
}

//文章
type Article struct {
	ArticleId      int
	ArticleTitle   string
	ArticleContext string
	ArticleAuthor  string
	ArticleTime    int64
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
	Title    string `json:"title"`
	Aboutimg string `json:"aboutimg"`
	Info1    string `json:"info1"`
	Info2    string `json:"info2"`
	Info3    string `json:"info3"`
	Btn1     string `json:"btn1"`
	Btn2     string `json:"btn2"`
	Btnx2    string `json:"btnx2"`
	Btnf3    string `json:"btnf3"`
	Infox1   string `json:"infox1"`
	Infox2   string `json:"infox2"`
	Infox3   string `json:"infox3"`
	Infox4   string `json:"infox4"`
	Infox5   string `json:"infox5"`
	Infox6   string `json:"infox6"`
	Infof1   string `json:"infof1"`
	Infof2   string `json:"infof2"`
	Infof3   string `json:"infof3"`
	Infof4   string `json:"infof4"`
	Infod1   string `json:"infod1"`
	Infod2   string `json:"infod2"`
	Infod3   string `json:"infod3"`
	Infod4   string `json:"infod4"`
	Infod5   string `json:"infod5"`
}

type User struct {
	Id       int
	Username string
	Password string
	Qq       string
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

func (User) TableName() string {
	return "user"
}
