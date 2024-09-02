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
}
