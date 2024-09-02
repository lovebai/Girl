package dao

import (
	"Girl/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

var Inx Index
var Mgr Manager

func init() {
	db, err := gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})

	if err != nil {
		panic("数据库打开失败")
	}

	Inx = &manager{db: db}
	Mgr = &manager{db: db}
	// Mgr = Inx

	//迁移 schema
	db.AutoMigrate(&model.Siteinfo{})

}
