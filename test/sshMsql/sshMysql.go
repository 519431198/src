package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
)

type Dialer struct {
	client *ssh.Client
}

type SSH struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
	Password string `json:"password"`
	KeyFile  string `json:"key"`
}

type MySQL struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (v *Dialer) Dial(address string) (net.Conn, error) {
	return v.client.Dial("tcp", address)
}

func (s *SSH) DialWithPassword() (*ssh.Client, error) {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", address, config)
}

func (s *SSH) DialWithKeyFile() (*ssh.Client, error) {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User:            s.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if k, err := ioutil.ReadFile(s.KeyFile); err != nil {
		return nil, err
	} else {
		signer, err := ssh.ParsePrivateKey(k)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	}
	return ssh.Dial("tcp", address, config)
}

func (m *MySQL) New() (db *gorm.DB, err error) {
	// 填写注册的mysql网络
	dsn := fmt.Sprintf("%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database)
	db, err = gorm.Open(sql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return
	}
	return
}

func main() {
	client := SSH{
		Host:     "10.0.0.10",
		User:     "root",
		Port:     22,
		Type:     "PASSWORD", // PASSWORD or KEY
		Password: "123",
	}
	my := MySQL{
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
	defer dial.Close()

	// 注册ssh代理
	mysql.RegisterDial("mysql+ssh", (&Dialer{client: dial}).Dial)

	db, err := my.New()
	if err != nil {
		log.Fatalf("mysql connect error: %s", err.Error())
		return
	}

	val := make(map[string]interface{})
	if err := db.Table("bill_hainan").Where("phone_number_x = 17119417195").Find(&val).Error; err != nil {
		log.Fatalf("mysql query error: %s", err.Error())
		return
	}
	fmt.Println(val["release_cause"])
}
