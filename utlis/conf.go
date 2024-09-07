package utlis

import (
	"Girl/model"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

const (
	FileName = "conf.ini"
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
		file, err := os.Create(FileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		defer file.Close() // 确保在函数结束时关闭文件

		content := fmt.Sprintf(`# possible values : release, debug, test
app_mode = debug 

[paths]
data = suibianba
path = Hello
		
[others]
salt = hello`)

		// 写入内容到文件
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

	}

	cfg, err := ini.Load(FileName)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	conf := model.Conf{}
	appMOde := cfg.Section("").Key("app_mode").String()
	data := cfg.Section("paths").Key("data").String()
	path := cfg.Section("paths").Key("path").String()
	salt := cfg.Section("others").Key("salt").String()
	if len(path) == 0 {
		appMOde = "release"
		data = "BoyandGirlDB"
		path = "Admin"
		salt = "hash123"

	}
	conf.AppMOde = appMOde
	conf.Data = data
	conf.Path = path
	conf.Salt = salt
	return conf
}

// 判断.db文件是否存在
func IsDbExists() bool {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		// fmt.Println("获取当前目录失败:", err)
		return false
	}

	// 查找以 _main.db 为结尾的文件
	var found bool
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件是否以 _main.db 结尾
		if !info.IsDir() && filepath.Base(info.Name()) == "_main.db" {
			found = true
			// fmt.Println("找到文件:", path)
		}
		return nil
	})

	if err != nil {
		// fmt.Println("遍历目录时出错:", err)
		return false
	}

	// if !found {
	// 	fmt.Println("当前目录下没有以 _main.db 为结尾的文件")
	// }
	return found
}
