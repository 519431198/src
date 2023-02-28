package main

import "github.com/wanghuiyt/ding"

func main() {
	//导入包 go get -u github.com/wanghuiyt/ding
	d := ding.Webhook{
		AccessToken: "f9f7e29c56a678facf1cafcc7bdd5263a4bef01e4335dc97436cb131c33bd439",
		Secret:      "SEC6d80d653ee3f110571b91c55320cb57f2ee89f9471317c46a6689570adb13c2a",
		//艾特所有人需要两个同时开启
		//EnableAt: true, // 开启艾特
		//AtAll:       true, // 艾特所有人
	}
	_ = d.SendMessage("周末天气真好啊!", "18326147303")
}
