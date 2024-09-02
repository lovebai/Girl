package dao

import "Girl/model"

// 前台获取站点信息
func (mgr *manager) GetSetting() model.Siteinfo {
	options := model.Siteinfo{}
	mgr.db.First(&options)
	return options

}

// 前台获取文章列表
func (mgr *manager) GetArticleList() []model.Article {
	articels := make([]model.Article, 10)
	mgr.db.Order("article_id desc").Find(&articels)
	return articels
}

// 根据Id获取文章
func (mgr *manager) GetArticle(id int) model.Article {
	article := model.Article{}
	mgr.db.First(&article, id)
	return article
}

// 获取留言列表
func (mgr *manager) GetLenvingList() []model.Lenving {
	Lenvings := make([]model.Lenving, 10)
	mgr.db.Order("lenving_id desc").Find(&Lenvings)
	return Lenvings
}

// 获取留言列表总数
func (mgr *manager) GetLenvingCount() int64 {
	var count int64
	Lenvings := model.Lenving{}
	mgr.db.Model(&Lenvings).Count(&count)
	return count
}

// 添加留言
func (mgr *manager) AddLenving(data model.Lenving) bool {
	e := mgr.db.Create(data)
	if e == nil {
		return false
	}
	return true
}

// 获取图片列表
func (mgr *manager) GetPhotoList() []model.Photo {
	photos := make([]model.Photo, 10)
	mgr.db.Order("img_id desc").Find(&photos)
	return photos
}

// 获取todolist
func (mgr *manager) GetTodoList() []model.TodoList {
	todolist := make([]model.TodoList, 10)
	mgr.db.Order("list_id desc").Find(&todolist)
	return todolist
}

// 获取关于
func (mgr *manager) GetAbout() model.About {
	about := model.About{}
	mgr.db.First(&about)
	return about
}
