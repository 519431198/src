package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db10 *gorm.DB
var db93 *gorm.DB

type judge struct {
	Id                    uint    `gorm:"column:id;size:20"`
	PhoneNumberA          *string `gorm:"column:phone_number_a;size:50"`
	PhoneNumberX          string  `gorm:"column:phone_number_x;size:50"`
	PhoneNumberB          *string `gorm:"column:phone_number_b;size:50"`
	PhoneNumberC          *string `gorm:"column:phone_number_c;size:50"`
	PhoneNumberY          *string `gorm:"column:phone_number_y;size:50"`
	Dgts                  *string `gorm:"column:dgts;size:50"`
	BindingId             *string `gorm:"column:binding_id;size:50"`
	CallType              *string `gorm:"column:call_type;size:50"`
	CallResult            *string `gorm:"column:call_result;size:50"`
	CallTime              string  `gorm:"column:call_time;size:50"`
	RingingTime           string  `gorm:"column:ringing_time;size:50"`
	StartTime             string  `gorm:"column:start_time;size:50"`
	ReleaseTime           string  `gorm:"column:release_time;size:50"`
	CallId                *string `gorm:"column:call_id;size:50"`
	ReleaseDirection      *string `gorm:"column:release_direction;size:50"`
	ReleaseCause          *string `gorm:"column:release_cause;size:50"`
	CallRecording         *string `gorm:"column:call_recording;size:50"`
	RecordingUrl          *string `gorm:"column:recording_url;size:256"`
	RecordingMode         *string `gorm:"column:recording_mode;size:50"`
	TransferPhoneNumber   *string `gorm:"column:transfer_phone_number;size:50"`
	TransferReason        *string `gorm:"column:transfer_reason;size:50"`
	Ability               *string `gorm:"column:ability;size:50"`
	CallRecognitionResult *string `gorm:"column:call_recognition_result;size:50"`
	AdditionalData        *string `gorm:"column:additional_data;size:50"`
}

type billHainan struct {
	Id                    uint    `gorm:"column:id;size:20"`
	PhoneNumberA          *string `gorm:"column:phone_number_a;size:50"`
	PhoneNumberX          string  `gorm:"column:phone_number_x;size:50"`
	PhoneNumberB          *string `gorm:"column:phone_number_b;size:50"`
	PhoneNumberC          *string `gorm:"column:phone_number_c;size:50"`
	PhoneNumberY          *string `gorm:"column:phone_number_y;size:50"`
	Dgts                  *string `gorm:"column:dgts;size:50"`
	BindingId             *string `gorm:"column:binding_id;size:50"`
	CallType              *string `gorm:"column:call_type;size:50"`
	CallResult            *string `gorm:"column:call_result;size:50"`
	CallTime              string  `gorm:"column:call_time;size:50"`
	RingingTime           string  `gorm:"column:ringing_time;size:50"`
	StartTime             string  `gorm:"column:start_time;size:50"`
	ReleaseTime           string  `gorm:"column:release_time;size:50"`
	CallId                *string `gorm:"column:call_id;size:50"`
	ReleaseDirection      *string `gorm:"column:release_direction;size:50"`
	ReleaseCause          *string `gorm:"column:release_cause;size:50"`
	CallRecording         *string `gorm:"column:call_recording;size:50"`
	RecordingUrl          *string `gorm:"column:recording_url;size:256"`
	RecordingMode         *string `gorm:"column:recording_mode;size:50"`
	TransferPhoneNumber   *string `gorm:"column:transfer_phone_number;size:50"`
	TransferReason        *string `gorm:"column:transfer_reason;size:50"`
	Ability               *string `gorm:"column:ability;size:50"`
	CallRecognitionResult *string `gorm:"column:call_recognition_result;size:50"`
	AdditionalData        *string `gorm:"column:additional_data;size:50"`
	Charge6               int     `gorm:"column:charge6;size:32"`
	Charge60              int     `gorm:"column:charge60;size:32"`
}

var mysqlLogger logger.Interface

// 本地数据库
func init() {
	username := "root"
	password := "123"
	host := "10.0.0.10"
	port := "3306"
	timeout := "10s"

	// 日志
	mysqlLogger = logger.Default.LogMode(logger.Info)
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
	db10 = dbClient
}

// 93 数据库
func init() {
	username := "dandan"
	password := "dwy123"
	host := "39.102.237.93"
	port := "12306"
	timeout := "10s"

	// 日志
	mysqlLogger = logger.Default.LogMode(logger.Info)
	// 连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/hn_bill?charset=utf8mb4&parseTime=True"+
		"&loc=Local&timeout=%s", username, password, host, port, timeout)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 是否跳过事务
		SkipDefaultTransaction: true,
		//Logger:                 logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	db93 = dbClient
}

// 初始化变量
func val() (insTableDay, insTableMonth, queryTable string, log10, log93 *gorm.DB) {
	now := time.Now()
	day := now.AddDate(0, 0, -1).Format("20060102")
	month := now.AddDate(0, 0, -1).Format("200601")
	// 插入的表
	insTableDay = "zjhn_" + day
	insTableMonth = "zjhn_" + month
	// 查询的表
	queryTable = "bill_hainan_" + day
	log10 = db10.Session(&gorm.Session{Logger: mysqlLogger})
	log93 = db93.Session(&gorm.Session{Logger: mysqlLogger})
	return insTableDay, insTableMonth, queryTable, log10, log93
}

// 插入数据
func execute(billList []billHainan, log10 *gorm.DB, insTableDay, insTableMonth string) {
	for _, res := range billList {
		start, _ := time.Parse("2006-01-02 15:04:05", res.StartTime)
		end, _ := time.Parse("2006-01-02 15:04:05", res.ReleaseTime)
		// 计算 end-start 的时间差,转换为 float 秒数,再转换为 int 类型
		diff := int(end.Sub(start).Seconds())
		// 计算 6 秒数
		if diff%6 == 0 {
			res.Charge6 = diff / 6
		} else {
			res.Charge6 = diff/6 + 1
		}
		// 计算分钟数
		if diff%60 == 0 {
			res.Charge60 = diff / 60
		} else {
			res.Charge60 = diff/60 + 1
		}
		// 将符合条件的数据插入表中
		log10.Table(insTableDay).Create(&res)
		log10.Table(insTableMonth).Create(&res)
	}
}

func main() {
	var count int64
	var limit = 1000
	var t judge
	var bill billHainan
	var billList []billHainan
	// 表 & db 链接
	insTableDay, insTableMonth, queryTable, log10, log93 := val()

	// 新建表日表和月表
	log10.Table(insTableDay).AutoMigrate(&bill)
	log10.Table(insTableMonth).AutoMigrate(&bill)

	// 查询符合条件的数量
	b := log93.Table(queryTable).Migrator().HasTable(&t)
	if !b {
	}

	log93.Table(queryTable).Where("start_time != release_time").Select("COUNT(1)").Scan(&count)
	// 原生语句,查询符合条件的数量
	//log93.Raw(`SELECT COUNT(1) FROM bill_hainan_20230805 WHERE start_time != release_time`).Scan(&count)
	//fmt.Println(count/1000 + 1)

	for page := 1; page <= int(count)/limit+1; page++ {
		// 第几页
		//fmt.Printf("第%d页", page)

		// 跳过已查询过的条数
		offset := (page - 1) * limit
		// 查询限定条件
		condition := "start_time!=release_time"
		// 查询并将结果映射到结构体切片中 billList
		log93.Table(queryTable).Limit(limit).Offset(offset).Where(condition).Find(&billList)
		// 插入数据
		execute(billList, log10, insTableDay, insTableMonth)
	}
}
