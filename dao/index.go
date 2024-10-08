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

// 获取指定数据的留言列表
func (mgr *manager) GetLenvingListLimit(limit int) []model.Lenving {
	Lenvings := make([]model.Lenving, 10)
	mgr.db.Order("lenving_id desc").Limit(limit).Find(&Lenvings)
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
	r := mgr.db.Create(data).RowsAffected
	if r == 0 {
		return false
	} else {
		return true
	}

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

// 根据性别查用户
func (mgr *manager) GetUserinfoBySex(sex int) model.User {
	//男为1，女为0
	user := model.User{}
	mgr.db.Where("sex = ?", sex).First(&user)
	return user
}

//根据ip地址查询
func (mgr *manager) GetBlackByIp(cip string) (int64, model.IpBlackList) {
	bl := model.IpBlackList{}
	row := mgr.db.Where("ip = ? ", cip).Find(&bl)
	return row.RowsAffected, bl
}
