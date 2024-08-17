package model

import "gorm.io/gorm"

//站点设置
type Setting struct {
	OptionId    int
	OptionName  string
	OptionValue string
}

func (Setting) TableName() string {
	return "setting"
}

// 后台用户
type User struct {
	gorm.Model

	Username string
	Password string
}
