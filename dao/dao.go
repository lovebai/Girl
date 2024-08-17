package dao

import (
	"Girl/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager interface {
	GetSetting() []model.Setting
	// GetSetting(options *model.Setting)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	db, err := gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})

	if err != nil {
		panic("数据库打开失败")
	}

	Mgr = &manager{db: db}

	//迁移 schema
	db.AutoMigrate(&model.Setting{})

}

func (mgr *manager) GetSetting() []model.Setting {
	options := []model.Setting{}
	mgr.db.Find(&options)
	return options

}
