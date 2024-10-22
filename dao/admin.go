package dao

import (
	"Girl/model"
	"Girl/utlis"
)

/** 查询 **/
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
	photo := model.Photo{}
	mgr.db.Model(&photo).Count(&count)
	return count
}

// 获取IP黑名单总数
func (mgr *manager) GetIpBlackListCountSum() int64 {
	var count int64
	bl := model.IpBlackList{}
	mgr.db.Model(&bl).Count(&count)
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

// 获取指定条数的留言列表
func (mgr *manager) GetLenvingListAdminLimit(sum int) []model.Lenving {
	Lenvings := make([]model.Lenving, 10)
	mgr.db.Order("lenving_id desc").Limit(sum).Find(&Lenvings)
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

// 获取ip黑名单列表
func (mgr *manager) GetIpBlackListAdmin() []model.IpBlackList {
	il := make([]model.IpBlackList, 10)
	mgr.db.Order("id desc").Find(&il)
	return il
}

// 获取关于数据
func (mgr *manager) GetAboutAdmin() model.About {
	about := model.About{}
	mgr.db.First(&about)
	return about
}

// 根据ID查询文章（点滴）
func (mgr *manager) GetArticleAdminByID(id int) model.Article {
	article := model.Article{}
	mgr.db.First(&article, id)
	return article
}

// 根据ID查询图片
func (mgr *manager) GetPhotoAdminByID(id int) model.Photo {
	ph := model.Photo{}
	mgr.db.First(&ph, id)
	return ph
}

// 根据ID查询清单
func (mgr *manager) GetTodoListAdminByID(id int) model.TodoList {
	t := model.TodoList{}
	mgr.db.First(&t, id)
	return t
}

/** 删除 **/
// admin dao根据id 删除留言
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

// admin dao根据id 删除点滴
func (mgr *manager) DeleteLittle(id int) int64 {
	del := model.Article{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

// admin dao根据id 删除图片
func (mgr *manager) DeletePhoto(id int) int64 {
	del := model.Photo{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

// admin dao根据id 删除清单
func (mgr *manager) DeleteTodoList(id int) int64 {
	del := model.TodoList{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

// admin dao根据id 删除IP
func (mgr *manager) DeleteIpBlackList(id int) int64 {
	del := model.IpBlackList{}
	rsu := mgr.db.Delete(&del, id)
	return rsu.RowsAffected
}

/** 新增 **/
// dao 新增点滴
func (mgr *manager) AddLittle(id int, name string, title string, text string) int64 {
	al := model.Article{
		ArticleId:      id,
		ArticleTitle:   title,
		ArticleContext: text,
		ArticleAuthor:  name,
		ArticleTime:    utlis.GetTimeUnix(),
	}
	rsu := mgr.db.Create(&al)
	return rsu.RowsAffected
}

// dao 新增图片
func (mgr *manager) AddPhoto(id int, t string, text string, url string) int64 {
	ap := model.Photo{
		ImgId:   id,
		ImgText: text,
		ImgUrl:  url,
		ImgTime: utlis.ConvertToTimestamp(t),
	}
	rsu := mgr.db.Create(&ap)
	return rsu.RowsAffected
}

// dao 新增清单
func (mgr *manager) AddTodoList(id int, status int, title string, imgUrl string) int64 {
	tl := model.TodoList{
		ListId:     id,
		ListStatus: status,
		ListText:   title,
		ListImgurl: imgUrl,
	}
	rsu := mgr.db.Create(&tl)
	return rsu.RowsAffected
}

// dao 新增IP黑名单
func (mgr *manager) AddIpBlackList(id int, ip string, commit string) int64 {
	bl := model.IpBlackList{
		Id:      id,
		Ip:      ip,
		Commit:  commit,
		Address: utlis.Get_ip_city(ip),
		Time:    utlis.GetTimeUnix(),
	}
	rsu := mgr.db.Create(&bl)
	return rsu.RowsAffected
}

/** 更新 **/
// dao 根据id更新点滴
func (mgr *manager) UpdateLittles(id int, title string, text string) int64 {
	up := model.Article{}
	mgr.db.First(&up, id)
	up.ArticleTitle = title
	up.ArticleContext = text
	rsu := mgr.db.Where("article_id = ?", id).Save(&up)
	return rsu.RowsAffected
}

// dao 根据id更新图片
func (mgr *manager) UpdatePhotos(ph model.Photo) int64 {
	p := model.Photo{}
	mgr.db.First(&p, ph.ImgId)
	p.ImgText = ph.ImgText
	p.ImgUrl = ph.ImgUrl
	p.ImgTime = ph.ImgTime
	rsu := mgr.db.Where("img_id = ?", ph.ImgId).Save(&p)
	return rsu.RowsAffected
}

// dao 根据id更新清单
func (mgr *manager) UpdateTodolists(tl model.TodoList) int64 {
	t := model.TodoList{}
	mgr.db.First(&t, tl.ListId)
	t.ListText = tl.ListText
	t.ListStatus = tl.ListStatus
	t.ListImgurl = tl.ListImgurl
	rsu := mgr.db.Where("list_id = ?", tl.ListId).Save(&t)
	return rsu.RowsAffected
}

// dao 根据id更新关于
func (mgr *manager) UpdateAbouts(ab model.About) int64 {
	// a := model.About{}
	// mgr.db.First(&a, 1)
	// a = ab
	// a.Id = 1
	// rsu := mgr.db.Where("id = ?", 1).Save(&a)
	ab.Id = 1
	rsu := mgr.db.Where("id = ?", 1).Save(&ab)
	return rsu.RowsAffected
}

// dao 根据id更新设置A
func (mgr *manager) UpdateSettingA(st model.SettingA) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.SiteName = st.SiteName
	info.SiteLogo = st.SiteLogo
	info.SiteDesc = st.SiteDesc
	info.Keywords = st.Keywords
	info.Discription = st.Discription
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// dao 根据id更新设置B
func (mgr *manager) UpdateSettingB(st model.SettingB) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.BoyName = st.BoyName
	info.GirlName = st.GirlName
	info.BoyImgUrl = st.BoyImgUrl
	info.GirlImgUrl = st.GirlImgUrl
	info.StartTime = utlis.ConvertToTimestamp(st.StartTime)
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// dao 根据id更新设置C
func (mgr *manager) UpdateSettingC(st model.SettingC) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.CardOne = st.CardOne
	info.CardOneDesc = st.CardOneDesc
	info.CardTwo = st.CardTwo
	info.CardTwoDesc = st.CardTwoDesc
	info.CardThree = st.CardThree
	info.CardThreeDesc = st.CardThreeDesc
	info.SiteIcp = st.SiteIcp
	info.SiteGaIcp = st.SiteGaIcp
	info.SiteCopyright = st.SiteCopyright
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// dao 根据id更新设置D
func (mgr *manager) UpdateSettingD(st model.SettingD) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.AvaterShowType = st.AvaterShowType
	info.LinkType = st.LinkType
	info.LinkApiUrl = st.LinkApiUrl
	info.LinkApiToken = st.LinkApiToken
	info.BgimgUrl = st.BgimgUrl
	info.SiteBlur = st.SiteBlur
	info.SitePajx = st.SitePajx
	info.SiteIcon = st.SiteIcon
	info.Watch = st.Watch
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// dao 根据id更新设置E
func (mgr *manager) UpdateSettingE(st model.SettingE) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.LenvingSum = st.LenvingSum
	info.BlockWord = st.BlockWord
	info.SiteRegexp = st.SiteRegexp
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// dao 根据id更新设置F
func (mgr *manager) UpdateSettingF(st model.SettingF) int64 {
	info := model.Siteinfo{}
	mgr.db.First(&info, 1)
	info.SiteHeadStyle = st.SiteHeadStyle
	info.SiteFootJs = st.SiteFootJs
	rsu := mgr.db.Where("option_id = ?", info.OptionId).Save(&info)
	return rsu.RowsAffected
}

// 根据用户名查用户
func (mgr *manager) GetUserinfoByName(name string) (int64, model.User) {
	user := model.User{}
	rsu := mgr.db.Where("username = ?", name).First(&user)
	return rsu.RowsAffected, user
}

// 查所有后台用户
func (mgr *manager) GetAllUserList() []model.User {
	user := make([]model.User, 10)
	mgr.db.Find(&user)
	return user
}

// 根据id更新用户信息
func (mgr *manager) UpdateUsername(id int, name string, qq string) int {
	user := model.User{}
	mgr.db.Where("id = ? ", id).First(&user)
	user.Username = name
	user.Qq = qq
	return int(mgr.db.Where("id = ? ", id).Save(&user).RowsAffected)
}

// 根据id更新用户信息
func (mgr *manager) UpdatePassword(id int, pass string, qq string) int {
	user := model.User{}
	mgr.db.Where("id = ? ", id).First(&user)
	user.Password = pass
	user.Qq = qq
	return int(mgr.db.Where("id = ? ", id).Save(&user).RowsAffected)
}

// 根据id更新用户信息
func (mgr *manager) UpdateQQ(id int, qq string) int {
	user := model.User{}
	mgr.db.Where("id = ? ", id).First(&user)
	user.Qq = qq
	return int(mgr.db.Where("id = ? ", id).Save(&user).RowsAffected)
}

// 根据id更新用户信息
func (mgr *manager) UpdateNamePassQQ(id int, name string, pass string, qq string) int {
	user := model.User{}
	mgr.db.Where("id = ? ", id).First(&user)
	user.Username = name
	user.Password = pass
	user.Qq = qq
	return int(mgr.db.Where("id = ? ", id).Save(&user).RowsAffected)
}
