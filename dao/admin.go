package dao

import (
	"Girl/model"
)

// 获取留言总数
func (mgr *manager) GetLenvingCountSum() int64 {
	var count int64
	Lenvings := model.Lenving{}
	mgr.db.Model(&Lenvings).Count(&count)
	return count
}

// 获取文章总数
func (mgr *manager) GetArticleCountSum() int64 {
	var count int64
	ariticles := model.Article{}
	mgr.db.Model(&ariticles).Count(&count)
	return count
}

// 获取清单总数
func (mgr *manager) GetTodoListCountSum() int64 {
	var count int64
	todolists := model.TodoList{}
	mgr.db.Model(&todolists).Count(&count)
	return count
}

// 获取图片总数
func (mgr *manager) GetPhotoCountSum() int64 {
	var count int64
	photo := model.TodoList{}
	mgr.db.Model(&photo).Count(&count)
	return count
}

// 获取站点信息
func (mgr *manager) GetSettingInfo() model.Siteinfo {
	options := model.Siteinfo{}
	mgr.db.First(&options)
	return options

}

// 获取留言列表
func (mgr *manager) GetLenvingListAdmin() []model.Lenving {
	Lenvings := make([]model.Lenving, 10)
	mgr.db.Order("lenving_id desc").Find(&Lenvings)
	return Lenvings
}

// 获取点滴列表
func (mgr *manager) GetArticleListAdmin() []model.Article {
	article := make([]model.Article, 10)
	mgr.db.Order("article_id desc").Find(&article)
	return article
}

// 获取相册列表
func (mgr *manager) GetPhotoListAdmin() []model.Photo {
	ph := make([]model.Photo, 10)
	mgr.db.Order("img_id desc").Find(&ph)
	return ph
}

// 获取清单列表
func (mgr *manager) GetTodoListAdmin() []model.TodoList {
	td := make([]model.TodoList, 10)
	mgr.db.Order("list_id desc").Find(&td)
	return td
}

// 获取关于数据
func (mgr *manager) GetAboutAdmin() model.About {
	about := model.About{}
	mgr.db.First(&about)
	return about
}

// 删除
func (mgr *manager) DeleteLenving(id int) int64 {
	del := model.Lenving{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
	// if status.RowsAffected == 0 {
	// 	return false
	// } else {
	// 	return true
	// }
}

func (mgr *manager) DeleteLittle(id int) int64 {
	del := model.Article{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

func (mgr *manager) DeletePhoto(id int) int64 {
	del := model.Photo{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

func (mgr *manager) DeleteTodoList(id int) int64 {
	del := model.TodoList{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}
