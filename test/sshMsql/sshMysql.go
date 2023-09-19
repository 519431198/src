package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/my/repo/test/sshMsql/untils"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"log"
)

func main() {
	// 创建 SSH 和 MySQL 的实例并设置连接配置
	client := untils.SSH{
		Host:     "10.0.0.10",
		User:     "root",
		Port:     22,
		Type:     "PASSWORD", // PASSWORD or KEY
		Password: "123",
	}
	my := untils.MySQL{
		Host:     "localhost",
		User:     "root",
		Password: "123",
		Port:     3306,
		Database: "bill",
	}

	var (
		dial *ssh.Client
		err  error
	)

	dial = sshLink(client, err)

	defer dial.Close()

	db := tunnel(dial, my)

	var table = "bill_ware_house"
	var condition = "phone_number_x = 18620458081"
	query(db, table, condition)
}

// 根据连接类型进行 SSH 连接
func sshLink(client untils.SSH, err error) (dial *ssh.Client) {
	// 根据连接类型进行 SSH 连接
	switch client.Type {
	case "KEY":
		dial, err = client.DialWithKeyFile()
	case "PASSWORD":
		dial, err = client.DialWithPassword()
	default:
		panic("unknown ssh type.")
	}
	if err != nil {
		log.Fatalf("ssh connect error: %s", err.Error())
		return
	}
	return dial
}

// 建立隧道,创建数据库链接
func tunnel(dial *ssh.Client, my untils.MySQL) (db *gorm.DB) {
	register := &untils.Dialer{
		Client: dial,
	}
	register.Client = dial

	// 注册 SSH 代理
	mysql.RegisterDial("mysql+ssh", register.Dial)

	// 创建 MySQL 数据库连接
	db, err := my.New()
	if err != nil {
		log.Fatalf("mysql connect error: %s", err.Error())
		return
	}
	return db
}

// sql查询
func query(db *gorm.DB, table, condition string) {
	type res struct {
		PhoneNumberA string `db:"phone_number_a"`
		PhoneNumberX string `db:"phone_number_x"`
		PhoneNumberB string `db:"phone_number_b"`
	}
	var Res res
	//val := make(map[string]interface{})
	fmt.Println("到这里了")
	// 查询 MySQL 数据库
	err := db.Debug().Table(table).Where(condition).Select("phone_number_a,phone_number_x,phone_number_b").Scan(&Res).Error
	//err := db.Debug().Table(table).Where(condition).Find(&val).Error
	if err != nil {
		log.Fatalf("mysql query error: %s", err.Error())
		return
	}
	// 打印查询结果中的 release_cause 字段
	fmt.Println(Res.PhoneNumberA)
}
