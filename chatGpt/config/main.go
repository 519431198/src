package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// 定义一个 config 全局变量-->Provinces
var config provinces

type provinces struct {
	NumberSegment string              `yaml:"numberSegment"`
	Provinces     map[string][]string `yaml:"province"`
}

func readConfig() {
	// 读取配置文件
	data, err := ioutil.ReadFile("/Users/wangyi/go/src/chatGpt/config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析YAML数据
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
func main() {
	// 读取配置文件
	readConfig()
	// 打印配置信息
	for provinceName, cities := range config.Provinces {
		fmt.Printf("%s:\n", provinceName)
		for _, city := range cities {
			fmt.Printf("%s ", city)
		}
		fmt.Println()
	}
}
