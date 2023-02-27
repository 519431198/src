package main

import "github.com/wanghuiyt/ding"

func main() {
	//导入包 go get -u github.com/wanghuiyt/ding
	d := ding.Webhook{
		AccessToken: "f9f7e29c56a678facf1cafcc7bdd5263a4bef01e4335dc97436cb131c33bd439",
		Secret:      "SEC6e004dca28307acaa0dbf36e601a9e3436b78c6d200f8ddd2b553ef9a409e854",
		EnableAt:    true, // 开启艾特
		//AtAll:       true, // 艾特所有人
	}
	_ = d.SendMessage("沙雕!", "18326147303")
}
