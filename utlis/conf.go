package utlis

import (
	"Girl/model"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

const (
	FileName = "data/conf.ini"
	Port     = "5200"
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
			os.Exit(1)
		}
		defer file.Close()

		// 要写入的内容
		content := `# 运行模式 : release, debug, test
app_mode = release

# Web服务端口
app_port = 5200

[paths]
# 数据库文件名
data = temp

# 后台登录地址
path = Admin

[others]
# 密码加密盐
salt = jiami
		`

		// 写入内容到文件
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		fmt.Printf("Tip: 请先修改conf.ini配置文件,修改完成重新启动即可。\n")
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

// 头部
func Header() {
	fmt.Println(` _       _ _              ______ _       _    _______              ______       _                   
| |     (_) |            / _____|_)     | |  (_______)            / _____)     | |                  
| |      _| |  _ ____   | /  ___ _  ____| |   _____ ___   ____   | /  ___  ___ | | ____ ____   ____ 
| |     | | | / ) _  )  | | (___) |/ ___) |  |  ___) _ \ / ___)  | | (___)/ _ \| |/ _  |  _ \ / _  |
| |_____| | |< ( (/ /   | \____/| | |   | |  | |  | |_| | |      | \____/| |_| | ( ( | | | | ( ( | |
|_______)_|_| \_)____)   \_____/|_|_|   |_|  |_|   \___/|_|       \_____/ \___/|_|\_||_|_| |_|\_|| |
                                                                                             (_____|
`)
}

// 提示
func Tip(port string) {
	fmt.Printf("Like Girl For Golang 程序已启动，请通过 %v 端口访问 ^_^\n", port)
}
