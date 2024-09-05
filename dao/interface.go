package dao

import "Girl/model"

//前台
type Index interface {
	GetSetting() model.Siteinfo
	GetArticleList() []model.Article
	GetArticle(id int) model.Article
	GetLenvingList() []model.Lenving
	GetLenvingCount() int64
	AddLenving(data model.Lenving) bool
	GetPhotoList() []model.Photo
	GetTodoList() []model.TodoList
	GetAbout() model.About
	GetUserinfoBySex(sex int) model.User
}

//后台
type Manager interface {
	//查
	GetLenvingCountSum() int64
	GetArticleCountSum() int64
	GetTodoListCountSum() int64
	GetPhotoCountSum() int64
	GetLenvingListAdmin() []model.Lenving
	GetSettingInfo() model.Siteinfo
	GetArticleListAdmin() []model.Article
	GetPhotoListAdmin() []model.Photo
	GetTodoListAdmin() []model.TodoList
	GetAboutAdmin() model.About

	GetAllUserList() []model.User

	GetArticleAdminByID(id int) model.Article
	GetPhotoAdminByID(id int) model.Photo
	GetTodoListAdminByID(id int) model.TodoList

	//删
	DeleteLenving(id int) int64
	DeleteLittle(id int) int64
	DeletePhoto(id int) int64
	DeleteTodoList(id int) int64

	//增
	AddLittle(id int, name string, title string, text string) int64
	AddPhoto(id int, t string, text string, url string) int64
	AddTodoList(id int, status int, title string, imgUrl string) int64

	//更新
	UpdateLittles(id int, title string, text string) int64
	UpdatePhotos(ph model.Photo) int64
	UpdateTodolists(tl model.TodoList) int64
	UpdateAbouts(ab model.About) int64

	UpdateSettingA(st model.SettingA) int64
	UpdateSettingB(st model.SettingB) int64
	UpdateSettingC(st model.SettingC) int64
	UpdateSettingD(st model.SettingD) int64
	UpdateSettingE(st model.SettingE) int64
	UpdateSettingF(st model.SettingF) int64

	//用户
	GetUserinfoByName(name string) (int64, model.User)
}
