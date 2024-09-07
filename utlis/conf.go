package utlis

import (
	"Girl/model"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

const (
	FileName = "conf.ini"
	Port     = "8888"
)

// 检查文件是否存在 存在则返回true
func GetFileExists() bool {
	if _, err := os.Stat(FileName); err != nil {
		return false
	}
	// fmt.Printf("File %s already exists.\n", FileName)
	return true
}

// 获取配置文件内容
func GetConfBody() model.Conf {
	if !GetFileExists() {
		fmt.Printf("Tip: %v 配置文件不存在，请先创建配置文件。\n", FileName)
		os.Exit(1)
	}

	cfg, err := ini.Load(FileName)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	conf := model.Conf{}
	appMode := cfg.Section("").Key("app_mode").String()
	appPort := cfg.Section("").Key("app_port").String()
	data := cfg.Section("paths").Key("data").String()
	path := cfg.Section("paths").Key("path").String()
	salt := cfg.Section("others").Key("salt").String()
	if len(path) == 0 {
		appMode = "release"
		appPort = Port
		data = "BoyandGirlDB"
		path = "Admin"
		salt = "hash123"

	}
	conf.AppMode = appMode
	conf.AppPort = appPort
	conf.Data = data
	conf.Path = path
	conf.Salt = salt
	return conf
}
