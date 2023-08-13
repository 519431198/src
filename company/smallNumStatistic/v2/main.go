package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"path/filepath"
	"time"
)

type BillWareHouse struct {
	ID                   string `gorm:"primaryKey;type:varchar(100);"`
	ServiceCode          string `gorm:"type:varchar(20)"`
	PhoneNumberA         string `gorm:"type:varchar(20)"`
	PhoneNumberAAreaCode string `gorm:"type:varchar(20)"`
	PhoneNumberAOperator string `gorm:"type:varchar(20)"`
	PhoneNumberAProvince string `gorm:"type:varchar(20)"`
	PhoneNumberACity     string `gorm:"type:varchar(20)"`
	PhoneNumberB         string `gorm:"type:varchar(50)"`
	PhoneNumberBAreaCode string `gorm:"type:varchar(20)"`
	PhoneNumberBOperator string `gorm:"type:varchar(20)"`
	PhoneNumberBProvince string `gorm:"type:varchar(20)"`
	PhoneNumberBCity     string `gorm:"type:varchar(20)"`
	PhoneNumberX         string `gorm:"type:varchar(20)"`
	PhoneNumberXAreaCode string `gorm:"type:varchar(20)"`
	PhoneNumberXOperator string `gorm:"type:varchar(20)"`
	PhoneNumberXProvince string `gorm:"type:varchar(20)"`
	PhoneNumberXCity     string `gorm:"type:varchar(20)"`
	PhoneNumberY         string `gorm:"type:varchar(20)"`
	PhoneNumberYAreaCode string `gorm:"type:varchar(20)"`
	PhoneNumberYOperator string `gorm:"type:varchar(20)"`
	PhoneNumberYProvince string `gorm:"type:varchar(20)"`
	PhoneNumberYCity     string `gorm:"type:varchar(20)"`
	ExtensionNumber      string `gorm:"type:varchar(20)"`
	BindingID            string `gorm:"type:varchar(50)"`
	CallID               string `gorm:"type:varchar(50)"`
	CallTime             string `gorm:"type:varchar(20)"`
	RingingTime          string `gorm:"type:varchar(20)"`
	StartTime            string `gorm:"type:varchar(20)"`
	ReleaseTime          string `gorm:"type:varchar(20)"`
	ReleaseDirection     string `gorm:"type:varchar(20)"`
	ReleaseCause         string `gorm:"type:varchar(20)"`
	CallRecording        string `gorm:"type:varchar(20)"`
	RecordingURL         string `gorm:"type:varchar(255)"`
	RecordingMode        string `gorm:"type:varchar(20)"`
	CallType             string `gorm:"type:varchar(20)"`
	CallResult           string `gorm:"type:varchar(20)"`
	TransferPhoneNumber  string `gorm:"type:varchar(20)"`
	TransferReason       string `gorm:"type:varchar(20)"`
	CallDuration         int    `gorm:"type:int(50)"`
	CustomerID           string `gorm:"type:varchar(50)"`
	CustomerName         string `gorm:"type:varchar(20)"`
	OpenID               string `gorm:"type:varchar(50)"`
	SMSContent           string `gorm:"type:varchar(255)"`
	Ability              string `gorm:"type:varchar(20)"`
	SMSSCount            int    `gorm:"type:int(50)"`
	Charge6              int    `gorm:"type:int(50)"`
	Charge60             int    `gorm:"type:int(50);index"`
}

// 查询结果
type scan struct {
	Min int64
	Sec int64
}

// Customer 解析配置文件
type Customer struct {
	Phones string `yaml:"phones"`
}

type Config struct {
	StartTime string              `yaml:"startTime"`
	EndTime   string              `yaml:"endTime"`
	Customers map[string]Customer `yaml:"customers"`
}

var db *gorm.DB

func main() {
	var bill *BillWareHouse
	var config *Config
	// 获取程序所在路径
	path := filePath()
	// 配置文件路径
	configPath := path + "config.yaml"
	configPath = "company/smallNumStatistic/v2/config.yaml"

	// 读取配置文件
	config = config.readConfig(configPath)
	// 选择统计方式
	var num int
	fmt.Print("1.统计昨天\n2.统计往期\n3.退出\n请选择(1-3): ")
	fmt.Scan(&num)
	switch num {
	case 1:
		resFilePath := path + time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		resFilePath = fmt.Sprintf("company/smallNumStatistic/v2/%s.csv", time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
		config.yesterday(bill, resFilePath)
	case 2:
		resFilePath := path + config.StartTime + config.EndTime
		resFilePath = fmt.Sprintf("company/smallNumStatistic/v2/%s-->%s.csv", config.StartTime, config.EndTime)
		config.history(bill, resFilePath)
	case 3:
		fmt.Print("再见!")
		os.Exit(0)
	default:
		fmt.Println("请输入(1/2)")
	}
}

// 获取程序所在路径
func filePath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取程序路径失败")
	}
	return path
}

// 统计昨日数据
func (config *Config) yesterday(bill *BillWareHouse, resFilePath string) {
	config.writeStr(resFilePath, "日期,客户,6秒数,分钟数\n")
	table := "bill_ware_house"
	for name, customer := range config.Customers {
		sec, min := config.querySummary(table, bill, customer)
		res := fmt.Sprintf("%s,%s,%d,%d\n", time.Now().AddDate(0, 0, -1).Format("2006-01-02"), name, sec, min)
		// 写入结果
		config.writeStr(resFilePath, res)
	}
}

// 统计往期数据
func (config *Config) history(bill *BillWareHouse, resFilePath string) {
	startTime, endTime := config.judgeTime()
	config.writeStr(resFilePath, "日期,客户,6秒数,分钟数\n")
	// 这里需要取到结束那一天,Before只能取到结束日期前一天,所以在结束日期上加一天
	for startTime.Before(endTime.AddDate(0, 0, 1)) {
		table := "bill_ware_house_" + startTime.Format("2006-01-02")
		for name, customer := range config.Customers {
			sec, min := config.querySummary(table, bill, customer)
			res := fmt.Sprintf("%s,%s,%d,%d\n", startTime.Format("2006-01-02"), name, sec, min)
			// 写入结果
			config.writeStr(resFilePath, res)
		}
		startTime = startTime.AddDate(0, 0, 1)
	}
}

// 公共查询逻辑
func (config *Config) querySummary(table string, bill *BillWareHouse, customer Customer) (int64, int64) {
	var sum scan
	// 查询条件
	query := fmt.Sprintf("phone_number_x IN (%s) AND phone_number_b != \"\" AND call_type < '200' "+
		"AND call_duration > '0';", customer.Phones)
	// 查询字段
	column := "SUM(charge60) as min,SUM(charge6) as sec"
	// 执行查询语句
	//db.Debug().Model(&bill).Table(table).Where(query).Select(column).Scan(&sum)
	db.Model(&bill).Table(table).Where(query).Select(column).Scan(&sum)
	return sum.Sec, sum.Min
}

// 初始化数据库
func init() {
	username := "dandan"
	password := "dwy123"
	host := "39.102.237.93"
	port := "12306"
	timeout := "10s"

	// 日志
	//mysqlLogger = logger.Default.LogMode(logger.Info)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/bill?charset=utf8mb4&parseTime=True"+
		"&loc=Local&timeout=%s", username, password, host, port, timeout)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		//Logger:                 logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	db = dbClient
}

// 读取配置文件
func (config *Config) readConfig(path string) *Config {
	configFile, err := os.Open(path)
	if err != nil {
		time.Sleep(time.Second * 3)
		log.Fatal("配置文件打开失败", err)
	}
	defer configFile.Close()
	// 创建一个文件读取对象(文件读取器)
	configNewReader := bufio.NewReader(configFile)
	// 创建一个 yaml 解码器
	yamlDecoder := yaml.NewDecoder(configNewReader)
	// 解析 yaml 内容,将结果反序列化到 config 结构体中
	err = yamlDecoder.Decode(&config)
	if err != nil {
		time.Sleep(time.Second * 3)
		log.Fatal("反序列化失败", err)
	}
	return config
}

// 判断时间格式是否正确,判断开始/结束时间是否正确
func (config *Config) judgeTime() (s, e time.Time) {
	now := time.Now().Format("2006-01-02")
	nowDay, _ := time.Parse("2006-01-02", now)
	// 起始时间
	s, err := time.Parse("2006-01-02", config.StartTime)
	if err != nil {
		log.Fatal("起始时间格式不正确")
	}
	// 结束时间
	e, _ = time.Parse("2006-01-02", config.EndTime)
	if err != nil {
		log.Fatal("结束时间,格式不正确")
	}

	// 判断结束日期与当前日期是否一致
	if e.Format("2006-01-02") != now {
		if !e.Before(nowDay) {
			log.Fatal("结束时间不得是当天时间及之后时间")
		}
	} else {
		log.Fatal("历史数据不统计当天数据,统计当天数据请选功能: 1")
	}
	// 判断开始时间是否在结束时间之前
	if !s.Before(e) {
		log.Fatal("开始时间不得晚于结束时间")
		return
	}
	return s, e
}

// 写如文件
func (config *Config) writeStr(resFilePath, resString string) {
	resFile, err := os.OpenFile(resFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		time.Sleep(time.Second * 3)
		log.Fatal("打开文件失败", err)
	}
	defer resFile.Close()
	writer := bufio.NewWriter(resFile)
	_, err = writer.WriteString(resString)
	if err != nil {
		log.Fatal("数据写入失败", err)
	}
	err = writer.Flush()
	if err != nil {
		time.Sleep(time.Second * 3)
		log.Fatal("刷新写入失败", err)
	}
}
