package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// 定义一个 config 全局变量--&gt;Provinces
var config provinces

// 定义 url 常量
const url = "https://shoujihao.uutool.cn/prefix/"

// 定义一个省结构体
type provinces struct {
	NumberSegment string              `yaml:"numberSegment"`
	Provinces     map[string][]string `yaml:"province"`
}

// 定义连接池
var dbPool *sql.DB

// 读取配置文件,反序列化到 config
func readConfig() {
	data, err := ioutil.ReadFile("/Users/wangyi/go/src/chatGpt/config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析YAML数据
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

// 初始化连接池
func initDB() {
	// 设置连接池的最大连接数和最大空闲连接数
	db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10:3306)/bill")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(500)

	// 检查连接池是否正常工作
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	dbPool = db
}

// 访问网页获取数据
func getWebsiteData(url string) []byte {
	// 跳过网页的 tls(https) 检查
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// 创建 http 链接
	client := &http.Client{Transport: tr}
	// 客户端请求网站
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	defer resp.Body.Close()
	websiteData, _ := ioutil.ReadAll(resp.Body)
	return websiteData
}

// 需要对网站返回数据进行判断,将数据进行处理
func judge(websiteData []byte) []string {
	// 匹配表达式
	ruler := fmt.Sprintf(">%s\\d{4}", config.NumberSegment)
	// 匹配规则
	pattern1 := regexp.MustCompile(ruler)
	// 匹配的结果
	numbers := pattern1.FindAllString(string(websiteData), -1)
	return numbers
}

// 正则匹配获取指定数据,拼接为 sql 语句
func province(res []string, pro, city string) {
	// 获取地市区号，邮政编码
	postalCode := getWebsiteData(url + strings.TrimLeft(res[1], ">"))
	// 匹配表达式
	ruler := ">\\d{3,6}<"
	// 匹配规则
	pattern2 := regexp.MustCompile(ruler)
	res2 := pattern2.FindAllString(string(postalCode), -1)
	// 区号
	cityNum := fmt.Sprintf("%s", strings.TrimLeft(strings.Trim(res2[len(res2)-2], "><"), "0"))
	// 邮编
	postNum := fmt.Sprintf("%s", strings.Trim(res2[len(res2)-1], "><"))
	//fmt.Println("号段为:")
	// sql 部分语句
	str := "INSERT INTO `bill`.`tb_mobile_number_section` (`id`,`number_section`,`operator_type`,`operator_name`,`virtual_operator_name`,`province`,`city`,`district`,`directly_city`,`card_type`,`area_code`,`post_code`,`is_display_redirnumber_info`) VALUES ("
	for _, v := range res {
		// sql 语句拼接
		sqlStr := fmt.Sprintf("\n%s'%s',%s,'非虚拟运营商','移动',NULL,'%s','%s市',NULL,0,'中国移动',%s,%s,0);", str, strings.TrimLeft(v, ">"), strings.TrimLeft(v, ">"), pro, city, cityNum, postNum)
		err := execSql(sqlStr)
		if err != nil {
			continue
		}
	}
}

// 执行 SQL 命令
func execSql(sqlOrder string) (err error) {
	// 获取连接池中的连接
	db := dbPool
	// 创建上下文，并指定最大超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 执行 SQL 命令
	_, err = db.ExecContext(ctx, sqlOrder)
	if err != nil {
		fmt.Printf("%s跳过\n", err)
	}

	return err
}

func main() {
	// 读取配置文件
	readConfig()

	// 初始化连接池
	initDB()

	// 打印配置信息
	for provinceName, cities := range config.Provinces {
		fmt.Printf("%s正在入库\n", provinceName)
		// cities 为城市集合,city 为单个城市
		for _, city := range cities {
			fmt.Println(city)
			numbers := getWebsiteData(url + city + config.NumberSegment)
			res := judge(numbers)
			if len(res) == 0 {
				continue
			}
			province(res, provinceName, city)
			fmt.Printf("%s入库完毕,等待 5 秒\n", city)
			//等待 5s 是因为频繁请求会禁止访问(大流量会禁止当天访问)
			time.Sleep(time.Second * 5)
		}
		fmt.Printf("%s入库完毕\n\n", provinceName)
	}
}
