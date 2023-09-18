package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

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

var db *gorm.DB

// 初始化数据库
func init() {
	//username := "dandan"
	//password := "dwy123"
	//host := "39.102.237.93"
	//port := "12306"
	//timeout := "10s"

	username := "root"
	password := "123"
	host := "10.0.0.10"
	port := "3306"
	timeout := "10s"

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
		panic("数据库连接失败" + err.Error())
	}

	db = dbClient
}

func main() {

	var num int
	fmt.Print("1.查询号码所属客户\n2.查询话单\n3.退出\n输入(1/2/3): ")
	for {
		fmt.Scan(&num)

		switch num {
		case 1:
			fmt.Println("查号码归属")
			fmt.Print("\n1.查询号码所属客户\n2.查询话单\n3.退出\n输入(1/2/3): ")
		case 2:
			query()
			fmt.Print("\n1.查询号码所属客户\n2.查询话单\n3.退出\n输入(1/2/3): ")
		case 3:
			os.Exit(1)
		default:
			fmt.Println("请输入 (1/2/3):")
		}
	}
}

func query() {
	for {
		var bill BillHn
		var num int64
		var phoneNumberX string
		var phoneNumberB string
		//17049650180 17765870601
		fmt.Print("输入号码信息,X 号码在前,B 号码在后以空格分隔(B 号未知部分用%代替): ")
		fmt.Scanln(&phoneNumberX, &phoneNumberB)
		condition := fmt.Sprintf("phone_number_x=%s and phone_number_x like %s", phoneNumberX, phoneNumberB)
		db.Debug().Model(&bill).Table("bill_hainan").Where(condition).Find("count(1)").Scan(&num)
		fmt.Println(num)
		if phoneNumberX == "1" {
			fmt.Println("退出程序!")
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}
	}
}
