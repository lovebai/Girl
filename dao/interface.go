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
}

//后台
type Manager interface {
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
}
