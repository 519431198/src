package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type RequestBody struct {
	Account string `json:"account"`
	//AreaCode  string `json:"areaCode"`
	Period    int    `json:"period"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
}

type ResponseBody struct {
	RetCode                        int                             `json:"retCode"`
	InfoReportCustomerLocationFees []InfoReportCustomerLocationFee `json:"infoReportCustomerLocationFees"`
}

type InfoReportCustomerLocationFee struct {
	BeginTime         int     `json:"beginTime"`
	EndTime           int     `json:"endTime"`
	AreaCode          string  `json:"areaCode"`
	AreaName          string  `json:"areaName"`
	Account           string  `json:"account"`
	AccountName       string  `json:"accountName"`
	AgentAccount      string  `json:"agentAccount"`
	CdrCount          float64 `json:"cdrCount"`
	TotalFee          float64 `json:"totalFee"`
	TotalTime         int     `json:"totalTime"`
	TotalSuiteFee     float64 `json:"totalSuiteFee"`
	TotalSuiteFeeTime int     `json:"totalSuiteFeeTime"`
}

func Handle(name, BeginTime, EndTime string, file *os.File) {
	url := "http://183.251.171.74:9090/external/server/GetReportCustomerLocationFee"
	requestBody := RequestBody{
		Account: name,
		//AreaCode:  "XY100002",
		Period:    1,
		BeginTime: BeginTime,
		EndTime:   EndTime,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("requestBody json 列化失败: ", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("req 创建请求失败: ", err)
		return
	}
	req.Header.Set("Content-Type", "text/html;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败: ", err)
		return
	}
	defer resp.Body.Close()
	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ResponseBodyBytes 读取响应失败", err)
		return
	}
	var responseBody ResponseBody
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		fmt.Println("responseBodyBytes 反序列化失败", err)
	}
	if responseBody.RetCode != 0 {
		fmt.Println("响应失败", responseBody.RetCode)
		return
	}

	for _, fee := range responseBody.InfoReportCustomerLocationFees {
		b, err := time.Parse("20060102", requestBody.BeginTime)
		if err != nil {
			fmt.Println("解析失败", err)
		}
		beginTime := b.Format("2006-01-02")
		_, err = file.Write([]byte(fmt.Sprintf("\n%v,%v,%v,%v,%v,%v,%v", beginTime, fee.AreaCode,
			fee.Account, fee.AccountName, fee.CdrCount, fmt.Sprintf("%0.3f", fee.TotalFee), fee.TotalTime/6)))
		//fmt.Printf("\n%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", beginTime,
		//	endTime, fee.AreaCode, fee.Account, fee.AccountName, fee.CdrCount,
		//	fmt.Sprintf("%0.3f", fee.TotalFee), fee.TotalTime, fee.TotalSuiteFee, fee.TotalSuiteFeeTime)
	}
}
