package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type logType struct {
	ServiceCode          string
	PhoneNumberA         string
	PhoneNumberAAreaCode string
	PhoneNumberAOperator string
	PhoneNumberAProvince string
	PhoneNumberACity     string
	PhoneNumberB         string
	PhoneNumberBAreaCode string
	PhoneNumberBOperator string
	PhoneNumberBProvince string
	PhoneNumberBCity     string
	PhoneNumberX         string
	PhoneNumberXAreaCode string
	PhoneNumberXOperator string
	PhoneNumberXProvince string
	PhoneNumberXCity     string
	PhoneNumberY         string
	PhoneNumberYAreaCode string
	PhoneNumberYOperator string
	PhoneNumberYProvince string
	PhoneNumberYCity     string
	ExtensionNumber      string
	BindingId            string
	CallId               string
	CallTime             string
	RingingTime          string
	StartTime            string
	ReleaseTime          string
	ReleaseDirection     string
	ReleaseCause         string
	CallRecording        string
	RecordingUrl         string
	RecordingMode        string
	CallType             string
	CallResult           string
	TransferPhoneNumber  string
	TransferReason       string
	CallDuration         int
	CustomerId           string
	CustomerName         string
	OpenId               string
	SmsContent           string
	Ability              string
	SmsCount             int
	Charge6              int
	Charge60             int
}

//打开文件,创建并返回文件句柄
func openFile(filePath string) (file *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err =", err)
		return
	}
	return file
}

//打开一个文件,如果不存在则创建
func writeFile(filePath string) (file *os.File) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
	}
	return file
}

//读取文件的每一行
func writeChannel(file *os.File, byteChan chan []byte) {
	lineReader := bufio.NewReader(file)
	//for i := 0; i < 10000; i++ {
	for {
		line1, err := lineReader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		byteChan <- []byte(line1)
	}
	close(byteChan)
}

func readfile(byteChan chan []byte, file2 *os.File, exitChan chan bool) {
	//for i := 0; i < 10000; i++ {
	for {
		line2, ok := <-byteChan
		if !ok {
			break
		}
		//fmt.Println(string(line2))
		var data logType
		err := json.Unmarshal(line2, &data)
		if err != nil {
			fmt.Println("json.Unmarshal err = ", err)
			break
		}
		if data.Charge60 == 0 && data.OpenId != "403258633c1d4b5a9dc5a4e49b313b19" {
			continue
		}
		str := fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", data.PhoneNumberA, data.PhoneNumberX, data.PhoneNumberB, data.StartTime, data.ReleaseTime, data.Charge60)
		writer := bufio.NewWriter(file2)
		writer.WriteString(str)
		writer.Flush()
		fmt.Println(str)
		//fmt.Printf("%v,%v,%v,%v,%v,%v\n", data.PhoneNumberA, data.PhoneNumberX, data.PhoneNumberB, data.StartTime, data.ReleaseTime, data.Charge60)
	}
	//fmt.Println("取不到数据,退出!")
	exitChan <- true
}

func main() {
	var num = 4
	var filePath1 = `/Users/wangyi/json.log`
	var filePath2 = `/Users/wangyi/403258633c1d4b5a9dc5a4e49b313b19.log`
	var byteChan = make(chan []byte, 100)
	var exitChan = make(chan bool, 4)
	file1 := openFile(filePath1)
	file2 := writeFile(filePath2)
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {

		}
	}(file1)
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {

		}
	}(file2)
	go writeChannel(file1, byteChan)
	for i := 0; i < num; i++ {
		go readfile(byteChan, file2, exitChan)
	}
	for i := 0; i < num; i++ {
		<-exitChan
	}
	close(exitChan)
}
