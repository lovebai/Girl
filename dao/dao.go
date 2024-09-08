package dao

import (
	"Girl/model"
	"Girl/utlis"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

var Ist Install
var Inx Index
var Mgr Manager

// 初始化数据库
func init() {
	dataFilepath := "data/" + utlis.GetConfBody().Data + "_main.db"
	db, err := gorm.Open(sqlite.Open(dataFilepath), &gorm.Config{})

	if err != nil {
		panic("数据库打开失败")
	}

	Ist = &manager{db: db}
	Inx = &manager{db: db}
	Mgr = &manager{db: db}
	// Mgr = Inx

	//迁移 schema
	db.AutoMigrate(&model.Siteinfo{}, &model.About{}, &model.Lenving{}, &model.Article{}, &model.Photo{}, &model.TodoList{}, &model.IpBlackList{}, &model.User{})

}
