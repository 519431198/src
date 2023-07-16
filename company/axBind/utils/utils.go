package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// GetWebsiteData 访问网页获取数据
func GetWebsiteData(url string) []byte {
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

// Code 获取手机区号
func Code(str string) string {

	rule := `归属地区号.*\n(.*)`
	reg := regexp.MustCompile(rule)
	match := reg.FindStringSubmatch(str)

	reg = regexp.MustCompile("[^0-9]+")
	return reg.ReplaceAllString(match[1], "")
}
