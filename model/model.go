package model

import "gorm.io/gorm"

// 后台用户模型
type User struct {
	gorm.Model

	Username string
	Password string
}

//站点设置

type Setting struct {
	gorm.Model
}
