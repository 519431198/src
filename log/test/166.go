package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// 定义 log 结构体
type logs struct {
	Timestamp string `json:"@timestamp"`
	Fields    fields `json:"@fields"`
}

// nginx 日志字段信息
type fields struct {
	RemoteAddr    string `json:"remote_addr"`
	Status        string `json:"status"`
	BodyBytesSent string `json:"body_bytes_sent"`
	RequestTime   string `json:"request_time"`
	Request       string `json:"request"`
	RequestMethod string `json:"request_method"`
	HttpUserAgent string `json:"http_user_agent"`
	RequestBody   string `json:"request_body"`
}

// 日志中请求信息
type requestBody struct {
	PhoneNumberA        string      `json:"phoneNumberA"`
	PhoneNumberX        string      `json:"phoneNumberX"`
	PhoneNumberB        string      `json:"phoneNumberB"`
	PhoneNumberC        string      `json:"phoneNumberC"`
	Dgts                string      `json:"dgts"`
	BindingId           string      `json:"bindingId"`
	CallType            string      `json:"callType"`
	CallTime            string      `json:"callTime"`
	RingingTime         string      `json:"ringingTime"`
	StartTime           string      `json:"startTime"`
	ReleaseTime         string      `json:"releaseTime"`
	CallId              string      `json:"callId"`
	ReleaseDirection    string      `json:"releaseDirection"`
	ReleaseCause        string      `json:"releaseCause"`
	CallRecording       string      `json:"callRecording"`
	RecordingUrl        interface{} `json:"recordingUrl"`
	RecordingMode       string      `json:"recordingMode"`
	TransferPhoneNumber string      `json:"transferPhoneNumber"`
	TransferReason      string      `json:"transferReason"`
	CallResult          string      `json:"callResult"`
	PhoneNumberY        string      `json:"phoneNumberY"`
	Ability             string      `json:"ability"`
	AdditionalData      string      `json:"additionalData"`
	InvokeId            string      `json:"invokeId"`
}

func main() {
	read := "/Users/wangyi/access_2023-07-12.log"
	write := "/Users/wangyi/zhejiang.csv"
	logFile, resFile := readWrite(read, write)
	defer logFile.Close()
	defer resFile.Close()

	// 创建一个通道，用于存储文件内容
	lines := make(chan string)
	var mu sync.Mutex
	companyResults := make(map[string]int64)
	// 创建一个等待组
	wg := sync.WaitGroup{}

	// 启动 4 个协程并发处理文件内容
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			// 创建结构体实例
			var logs logs
			var requestBody requestBody

			var xzTime, hzwlTime, mlTime, ty, tyzn int64

			for line := range lines {
				// 搜索关键字
				if strings.Contains(line, "finish") {
					err := json.Unmarshal([]byte(line), &logs)
					if err != nil {
						fmt.Println(errors.New("logs 日志反序列化错误"))
						continue
					}
					err = json.Unmarshal([]byte(logs.Fields.RequestBody), &requestBody)
					if err != nil {
						fmt.Println(errors.New("requestBody 日志反序列化错误"))
						continue
					}

					// 计算分钟数
					feeTime := result(&requestBody)
					if feeTime == 0 {
						continue
					}

					// 根据客户特征,匹配符合的日志
					customer, err := match(&logs)
					if err != nil {
						return
					}

					//fmt.Printf("%s,%d\n", customer, feeTime)
					// 将四条协程结果汇总
					switch customer {
					case "hzwl_axb_finish", "hzwl_ax_finish":
						hzwlTime += feeTime
					case "xz_axb_finish":
						xzTime += feeTime
					case "ml_axb_finish":
						mlTime += feeTime
					case "ty_axb_finish":
						ty += feeTime
					case "tyzn_axb_finish":
						tyzn += feeTime
					}
				}
			}

			// 输出每个协程处理的结果
			//fmt.Printf("协程 %d：杭州微聊:%d\n兴州智能:%d\n浙江美连:%d\n北京天眼:%d\n"+
			//	"北京天眼智能:%d\n", i+1, hzwlTime, xzTime, mlTime, ty, tyzn)
			mu.Lock()
			companyResults["杭州微聊"] += hzwlTime
			companyResults["兴州智能"] += xzTime
			companyResults["浙江美连"] += mlTime
			companyResults["北京天眼"] += ty
			companyResults["北京天眼智能"] += tyzn
			mu.Unlock()
			wg.Done()
		}()
	}

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(logFile)

	// 将文件内容写入通道中
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	// 关闭通道
	close(lines)

	// 等待所有协程处理完毕
	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	writeRes(resFile, companyResults)
}

func match(logs *logs) (string, error) {
	pattern := "POST /(.*) HTTP(.*)"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}
	// 匹配字符串
	match := re.FindStringSubmatch(logs.Fields.Request)
	customer := ""
	if len(match) > 1 {
		customer = match[1]
	} else {
		fmt.Println("no match found")
	}
	return customer, err
}

// 计算结果
func result(res *requestBody) int64 {
	s, _ := time.Parse("2006-01-02 15:04:05", res.StartTime)
	r, _ := time.Parse("2006-01-02 15:04:05", res.ReleaseTime)
	starTime := s.Unix()
	releaseTime := r.Unix()
	feeTime := releaseTime - starTime
	if feeTime == 0 {
		return feeTime
	}
	if feeTime%60 != 0 {
		feeTime = feeTime/60 + 1
	}
	if feeTime%60 == 0 {
		feeTime = feeTime / 60
	}
	return feeTime
}

// 打开日志文件文件及写入结果文件
func readWrite(read, write string) (*os.File, *os.File) {
	readFile, err := os.Open(read)
	if err != nil {
		log.Fatalln(err)
	}
	writeFile, err := os.OpenFile(write, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return readFile, writeFile
}

// 循环读取 map,将结果写入文件
func writeRes(w *os.File, companyResults map[string]int64) {
	var res string
	for i, v := range companyResults {
		res += fmt.Sprintf("%s: %d\n", i, v)
	}
	writer := bufio.NewWriter(w)
	_, err := writer.WriteString(res)
	if err != nil {
		fmt.Println(errors.New("写入结果到文件错误"))
		return
	}
	err = writer.Flush()
}
