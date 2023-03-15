package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Global   global   `yaml:"global"`
	Database database `yaml:"database"`
	Cache    cache    `yaml:"cache"`
	Service  service  `yaml:"service"`
}

type global struct {
	Debug    bool   `yaml:"debug"`
	LogLevel string `yaml:"log_level"`
	ApiKey   string `yaml:"api_key"`
}

type database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type cache struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type service struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

// 定义一个全局变量
var ConfigData *config

// 读取配置文件
func loadConfig(fileName string) error {
	config := new(config)

	// 获取程序所在目录 os.Executable(): 获取可执行文件的的绝对路径包含文件名
	//exePath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(exePath)
	// 获取文件所在目录路径
	//exeDir := filepath.Dir(exePath)
	//fmt.Println(exeDir)

	// filepath.Abs(filepath.Dir(os.Args[0])) 返回的是 os.Args[0] 的上级目录的绝对路径,不包含文件信息
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exeDir)
	filePath := filepath.Join(exeDir, fileName)
	data, err := ioutil.ReadFile(filePath)

	// 反序列化yaml格式内容到结构体中
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	// 将解析后的配置文件数据存入全局变量中
	ConfigData = config
	return nil
}

func main() {
	err := loadConfig("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 输出配置文件内容
	fmt.Printf("%+v\n", *ConfigData)
}
