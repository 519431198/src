package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"time"
)

var db *gorm.DB

type BillHn struct {
	Id                    uint   `gorm:"column:id;size:20"`
	PhoneNumberA          string `gorm:"column:phone_number_a;size:50"`
	PhoneNumberX          string `gorm:"column:phone_number_x;size:50"`
	PhoneNumberB          string `gorm:"column:phone_number_b;size:50"`
	PhoneNumberC          string `gorm:"column:phone_number_c;size:50"`
	PhoneNumberY          string `gorm:"column:phone_number_y;size:50"`
	Dgts                  string `gorm:"column:dgts;size:50"`
	BindingId             string `gorm:"column:binding_id;size:50"`
	CallType              string `gorm:"column:call_type;size:50"`
	CallResult            string `gorm:"column:call_result;size:50"`
	CallTime              string `gorm:"column:call_time;size:50"`
	RingingTime           string `gorm:"column:ringing_time;size:50"`
	StartTime             string `gorm:"column:start_time;size:50"`
	ReleaseTime           string `gorm:"column:release_time;size:50"`
	CallId                string `gorm:"column:call_id;size:50"`
	ReleaseDirection      string `gorm:"column:release_direction;size:50"`
	ReleaseCause          string `gorm:"column:release_cause;size:50"`
	CallRecording         string `gorm:"column:call_recording;size:50"`
	RecordingUrl          string `gorm:"column:recording_url;size:256"`
	RecordingMode         string `gorm:"column:recording_mode;size:50"`
	TransferPhoneNumber   string `gorm:"column:transfer_phone_number;size:50"`
	TransferReason        string `gorm:"column:transfer_reason;size:50"`
	Ability               string `gorm:"column:ability;size:50"`
	CallRecognitionResult string `gorm:"column:call_recognition_result;size:50"`
	AdditionalData        string `gorm:"column:additional_data;size:50"`
}

// Customer 解析配置文件
type Customer struct {
	Phones string
}
type Customers struct {
	Customers map[string]Customer
}

// 查询结果
type scan struct {
	Min int64   `db:"min"`
	Sec int64   `db:"sec"`
	Tot int64   `db:"tot"`
	Suc int64   `db:"suc"`
	Per float32 `db:"per"`
}

//var mysqlLogger logger.Interface

// 初始化数据库
func init() {
	username := "dandan"
	password := "dwy123"
	host := "39.102.237.93"
	port := "12306"
	timeout := "10s"

	// 日志
	//mysqlLogger = logger.Default.LogMode(logger.Info)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/hn_bill?charset=utf8mb4&parseTime=True"+
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

func main() {
	now := time.Now()
	//end := "2023-08-01 18:00:00"
	// 获取前一天时间
	oldTime := now.AddDate(0, 0, 0).Format("15:04:05")

	//// 获取程序所在路径
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	// 获取当前路径下文件
	resDataPath := filepath.Join(fmt.Sprintf("%s/%s.csv", binPath, "nowData"))

	//resDataPath := "D:\\go\\src\\company\\hainan\\" + oldTime + ".csv"
	//resDataPath := "/Users/wangyi/goProject/project_1/src/company/zhejiang/nowData.csv"

	resFile, err := os.OpenFile(resDataPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer resFile.Close()

	// 日志打印,测试用
	//db = db.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	var billHn BillHn

	// 创建结构体实例
	var config Customers
	// 获取当前路径下配置文件
	configPath := binPath + "/config.yaml"
	file, err := os.Open(configPath)
	// 打开配置文件,创建文件句柄
	//file, err := os.Open("/Users/wangyi/goProject/project_1/src/company/zhejiang/config.yaml")

	if err != nil {
		fmt.Println("打开文件失败: ", err)
	}
	defer file.Close()
	// 创建读取器
	reader := bufio.NewReader(file)
	// 读取配置文件内容
	decoder := yaml.NewDecoder(reader)
	// 反序列化到结构体
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("解析配置文件失败", err)
	}
	// 拼接定义前一天表名
	//table := "bill_hainan_" + now.AddDate(0, 0, -1).Format("20060102")
	table := "bill_hainan"
	for name, customer := range config.Customers {
		// 选择列或对列的操作
		column := "sum(ceil( time_to_sec( timediff(release_time, start_time) ) / 60 )) as min," +
			"SUM(ceil( time_to_sec( timediff(release_time, start_time) ) / 6 )) as sec,COUNT(1) as tot," +
			"COUNT(CASE WHEN start_time != release_time THEN 1 END) AS suc," +
			"FORMAT(100 * COUNT(CASE WHEN start_time != release_time THEN 1 END)/ COUNT(1),2) AS per"
		// 定义 where 条件
		query := fmt.Sprintf("call_type < '200' AND phone_number_b != \"\" AND phone_number_x IN (%s)", customer.Phones)
		// 创建结果实例
		var count scan
		// 查询数据库
		db.Model(&billHn).Table(table).Where(query).Select(column).Scan(&count)
		//fmt.Printf("%s,%s,%d\n", oldTime, name, count)
		_, err = resFile.WriteString(fmt.Sprintf("%s,%s,%d,%d,%.2f%%,%d,%d\n", oldTime, name, count.Tot, count.Suc, count.Per, count.Min, count.Sec))
		if err != nil {
			fmt.Println("数据写入失败", err)
		}
		err = resFile.Sync()
		if err != nil {
			fmt.Println("刷新写入失败", err)
		}
	}
}
