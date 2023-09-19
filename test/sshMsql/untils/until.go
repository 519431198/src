package untils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"net"
)

// Dialer 结构体实现了 net.Dialer 接口，用于在 SSH 连接中进行网络连接。
type Dialer struct {
	Client *ssh.Client
}

// SSH 结构体包含 SSH 连接所需的配置信息。
type SSH struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
	Password string `json:"password"`
	KeyFile  string `json:"key"`
}

// MySQL 结构体包含 MySQL 数据库的配置信息。
type MySQL struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// Dial 方法在 SSH 连接中进行网络连接。
func (v *Dialer) Dial(address string) (net.Conn, error) {
	return v.Client.Dial("tcp", address)
}

// DialWithPassword 方法使用密码进行 SSH 连接。
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

// DialWithKeyFile 方法使用密钥文件进行 SSH 连接。
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

// New 方法创建一个新的 MySQL 数据库连接。
func (m *MySQL) New() (db *gorm.DB, err error) {
	// 填写注册的 mysql 网络
	dsn := fmt.Sprintf("%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database)
	db, err = gorm.Open(sql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	return
}
