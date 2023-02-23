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

//func openFile(filePath string) (lineReader *bufio.Reader) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		fmt.Println("open file err =", err)
//		return
//	}
//	lineReader = bufio.NewReader(file)
//	defer file.Close()
//	return lineReader
//}

// 读取文件的每一行
func readfile(filePath string, exitChan chan bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err =", err)
		return
	}
	defer file.Close()
	lineReader := bufio.NewReader(file)
	var byteChan chan []byte
	byteChan = make(chan []byte, 100)
	for i := 0; i <= 1; i++ {
		//for {
		line1, _, err := lineReader.ReadLine()
		//line1,err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		byteChan <- line1
		//byteChan <- []byte(line1)
		line2, _ := <-byteChan
		//fmt.Println(string(line2))
		var data logType
		err = json.Unmarshal(line2, &data)
		if err != nil {
			fmt.Println("json.Unmarshal err = ", err)
		}
		fmt.Printf("%v\n", data)
	}

	close(byteChan)
	exitChan <- true
	close(exitChan)
}
func main() {
	var filePath = `/Users/wangyi/json.log`
	var exitChan = make(chan bool, 4)
	//lineReader := openFile(filePath)
	readfile(filePath, exitChan)
	//time.Sleep(time.Second * 15)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
