package main

import (
	"crypto/tls"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"regexp"
	"time"
)

// 定义常量
const url = "https://shoujihao.uutool.cn/no/"

// 访问网页获取数据
func getWebsiteData(url string) []byte {
	// 跳过网页的 tls(https) 检查
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// 创建链接
	client := &http.Client{Transport: tr}
	// 客户端请求网站
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	defer resp.Body.Close()
	websiteData, _ := io.ReadAll(resp.Body)
	return websiteData
}

// 获取手机区号
func code(str string) string {

	rule := `归属地区号.*\n(.*)`
	reg := regexp.MustCompile(rule)
	match := reg.FindStringSubmatch(str)

	reg = regexp.MustCompile("[^0-9]+")
	return reg.ReplaceAllString(match[1], "")
}

func main() {
	phone := map[string]string{
		"13001140201": "13001140201",
		"13001140205": "13001140205",
		"13001140206": "13001140206",
		"13001140208": "13001140208",
		"13001140211": "13001140211",
		"13001140212": "13001140212",
	}

	for numX, _ := range phone {

		phoneNumberX := numX
		webSiteData := getWebsiteData(url + phoneNumberX)
		areaCode := code(string(webSiteData))
		fmt.Println(areaCode)
		time.Sleep(time.Second * 2)

	}
}
