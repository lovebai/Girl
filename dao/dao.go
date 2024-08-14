package dao

import (
	"Girl/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager interface {
}

type manager struct {
	db *gorm.DB
}

var Mgr manager

func init() {
	db, err := gorm.Open(sqlite.Open(""), &gorm.Config{})

	if err != nil {
		log.Fatal(&model.User{})
	}
	// Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})

}
