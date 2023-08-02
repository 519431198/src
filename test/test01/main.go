package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

// 初始化数据库
func init() {
	username := "dandan"
	password := "dwy123"
	host := "39.102.237.93"
	port := "12306"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/hn_bill?charset=utf8mb4&parseTime=True"+
		"&loc=Local&timeout=%s", username, password, host, port, timeout)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	db = dbClient
}

func main() {
	now := time.Now()
	// 获取前一天时间
	oldTime := now.AddDate(0, 0, -1).Format("2006-01-02")
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	configPath := binPath + "/config.yaml"

	resDataPath := filepath.Join(fmt.Sprintf("%s/%s.csv", binPath, oldTime))
	fmt.Println(configPath, resDataPath)
}
