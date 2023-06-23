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
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)

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
	// 创建 httpDemo01 链接
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

// 正则匹配获取指定数据,拼接为 sql 语句
func province(num, pro, city string, operator []string) {
	// 获取地市区号，邮政编码
	//postalCode := getWebsiteData(url + strings.TrimLeft(res[0], ">"))
	postalCode := getWebsiteData(url + num)
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
	str := "INSERT INTO `bill`.`tb_mobile_number_section` (`id`,`number_section`,`operator_type`,`operator_name`," +
		"`virtual_operator_name`,`province`,`city`,`district`,`directly_city`,`card_type`,`area_code`," +
		"`post_code`,`is_display_redirnumber_info`) VALUES ("
	// sql 语句拼接
	sqlStr := fmt.Sprintf("\n%s%s,%s,'%s','%s',NULL,'%s省','%s市',NULL,0,'中国%s',%s,%s,0);",
		str, num, num, operator[1], operator[0], pro, city, operator[0], cityNum, postNum)
	fmt.Println(sqlStr)
	err := execSql(sqlStr)
	if err != nil {
		fmt.Println(err)
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

func address(num string) (provinceName, cityName, err string) {
	websiteData := getWebsiteData(url + num)
	// 匹配指定信息
	re := regexp.MustCompile("<td>号段你归属地</td>\\s+<td>\\s*([\\s\\S]*?)\\s*</td>|404")
	matches := re.FindStringSubmatch(string(websiteData))
	if matches[0] == "404" {
		err = fmt.Sprintf("没有%s号段信息", num)
	} else {
		// 将多余部分替换删除
		replace := strings.NewReplacer(" ", "", ".*>", "", "号段你归属地", "", "<td>", "", "</td>", "", "<a href=\"/prefix/", "")
		s := replace.Replace(matches[0])
		// 取出号码归属地省,市
		re = regexp.MustCompile(".*1")
		result := re.FindAllString(s, 2)
		provinceName = strings.Replace(result[0], "1", "", -1)
		cityName = strings.Replace(result[1], "1", "", -1)
	}
	return provinceName, cityName, err
}

func main() {
	// 读取配置文件
	readConfig()
	// 初始化连接池
	initDB()
	operator := []string{"电信", "虚拟运营商"}
	num := []string{"1655726"}
	for _, num := range num {
		provinceName, cityName, err := address(num)
		if err != "" {
			continue
		}
		province(num, provinceName, cityName, operator)
	}
}
