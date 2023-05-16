package main

import (
	"fmt"
	"github.com/my/repo/post/utils"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type config struct {
	Names []string `json:"names"`
}

func (config *config) name() (names []string) {
	data, err := os.ReadFile("/code/bill_total/name.yaml")
	if err != nil {
		fmt.Println("配置文件读取失败", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("配置文件反序列化失败,err", err)
	}
	names = config.Names
	return names
}

func openFile() (file *os.File) {
	file, err := os.OpenFile("/code/bill_total/vos.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println()
	}
	return file
}

func main() {
	file := openFile()
	//_, err := file.Write([]byte("beginTime(起始时间),endTime(终止时间),areaCode(地区前缀),account(账户号码),accountName(账户名称),cdrCount(话单总计),totalFee(费用总计),totalTime(计费时长总计--秒),totalSuiteFee(套餐费用总计),totalSuiteFeeTime(套餐赠送时长总计)"))
	_, err := file.Write([]byte("起始时间,终止时间,地区前缀,账户号码,账户名称,话单总计,费用总计,计费时长总计--秒,套餐费用总计,套餐赠送时长总计"))
	if err != nil {
		fmt.Println("写入标题失败")
	}
	var config config
	names := config.name()
	// 获取当前时间
	now := time.Now()
	EndTime := now.Format("20060102")
	// 获取前一天日期
	b := now.AddDate(0, 0, -1)
	BeginTime := b.Format("20060102")
	// 循环读取 A 路账户
	for _, name := range names {
		utils.Handle(name, BeginTime, EndTime, file)
	}
	//fmt.Println("写入完成!")
}
