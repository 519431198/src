package main

import (
	"github.com/wanghuiyt/ding"
	"log"
)

func main() {
	//导入包 go get -u github.com/wanghuiyt/ding
	d := ding.Webhook{
		AccessToken: "0b8e5c9299cc7ca77cc863381e8b5949c648a4804ae3993b046e3ee4ab33b70d",
		Secret:      "SEC4b3633a968fe18595a855597c2be8315917e26b08b2eafe2f82f4d40314271eb",
		//艾特所有人需要两个同时开启
		EnableAt: true,  // 开启艾特
		AtAll:    false, // 艾特所有人
	}
	err := d.SendMessage("周末天气真好啊!", "18326147303")
	if err != nil {
		log.Fatalln(err)
	}
}
