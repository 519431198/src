package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Hd struct {
	ServiceCode          string `json:"serviceCode"`
	PhoneNumberA         string `json:"phoneNumberA"`
	PhoneNumberAAreaCode string `json:"phoneNumberAAreaCode"`
	PhoneNumberAOperator string `json:"phoneNumberAOperator"`
	PhoneNumberAProvince string `json:"phoneNumberAProvince"`
	PhoneNumberACity     string `json:"phoneNumberACity"`
	PhoneNumberB         string `json:"phoneNumberB"`
	PhoneNumberBAreaCode string `json:"phoneNumberBAreaCode"`
	PhoneNumberBOperator string `json:"phoneNumberBOperator"`
	PhoneNumberBProvince string `json:"phoneNumberBProvince"`
	PhoneNumberBCity     string `json:"phoneNumberBCity"`
	PhoneNumberX         string `json:"phoneNumberX"`
	PhoneNumberXAreaCode string `json:"phoneNumberXAreaCode"`
	PhoneNumberXOperator string `json:"phoneNumberXOperator"`
	PhoneNumberXProvince string `json:"phoneNumberXProvince"`
	PhoneNumberXCity     string `json:"phoneNumberXCity"`
	PhoneNumberY         string `json:"phoneNumberY"`
	PhoneNumberYAreaCode string `json:"phoneNumberYAreaCode"`
	PhoneNumberYOperator string `json:"phoneNumberYOperator"`
	PhoneNumberYProvince string `json:"phoneNumberYProvince"`
	PhoneNumberYCity     string `json:"phoneNumberYCity"`
	ExtensionNumber      string `json:"extensionNumber"`
	BindingId            string `json:"bindingId"`
	CallId               string `json:"callId"`
	CallTime             string `json:"callTime"`
	RingingTime          string `json:"ringingTime"`
	StartTime            string `json:"startTime"`
	ReleaseTime          string `json:"releaseTime"`
	ReleaseDirection     string `json:"releaseDirection"`
	ReleaseCause         string `json:"releaseCause"`
	CallRecording        string `json:"callRecording"`
	RecordingUrl         string `json:"recordingUrl"`
	RecordingMode        string `json:"recordingMode"`
	CallType             string `json:"callType"`
	CallResult           string `json:"callResult"`
	TransferPhoneNumber  string `json:"transferPhoneNumber"`
	TransferReason       string `json:"transferReason"`
	CallDuration         int    `json:"callDuration"`
	CustomerId           string `json:"customerId"`
	CustomerName         string `json:"customerName"`
	OpenId               string `json:"openId"`
	SmsContent           string `json:"smsContent"`
	Ability              string `json:"ability"`
	SmsCount             int    `json:"smsCount"`
	Charge6              int    `json:"charge6"`
	Charge60             int    `json:"charge60"`
}

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10:3306)/bill")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var jsonString = "{\"serviceCode\":\"2\",\"phoneNumberA\":\"17815610569\",\"phoneNumberAAreaCode\":null," +
		"\"phoneNumberAOperator\":\"移动\",\"phoneNumberAProvince\":null,\"phoneNumberACity\":null," +
		"\"phoneNumberB\":\"18826426529\",\"phoneNumberBAreaCode\":\"20\",\"phoneNumberBOperator\":\"移动\"," +
		"\"phoneNumberBProvince\":\"广东省\",\"phoneNumberBCity\":\"广州市\",\"phoneNumberX\":\"18620413738\"," +
		"\"phoneNumberXAreaCode\":\"20\",\"phoneNumberXOperator\":\"联通\",\"phoneNumberXProvince\":\"广东省\"," +
		"\"phoneNumberXCity\":\"广州市\",\"phoneNumberY\":\"18620413738\",\"phoneNumberYAreaCode\":\"20\"," +
		"\"phoneNumberYOperator\":\"联通\",\"phoneNumberYProvince\":\"广东省\",\"phoneNumberYCity\":\"广州市\"," +
		"\"extensionNumber\":null,\"bindingId\":\"\",\"callId\":\"4d63f43d2936043a\",\"callTime\":\"2023-02-21 11:40:25\"," +
		"\"ringingTime\":\"2023-02-21 11:40:28\",\"startTime\":\"2023-02-21 11:40:44\",\"releaseTime\":\"2023-02-21 11:42:06\"," +
		"\"releaseDirection\":\"1\",\"releaseCause\":\"31\",\"callRecording\":\"1\"," +
		"\"recordingUrl\":\"http://61.139.144.36:910/2023022111/RX16__17815610569_18620413738_18826426529_20230221114025_4d63f43d2936043a.mp3\"," +
		"\"recordingMode\":\"1\",\"callType\":\"110\",\"callResult\":\"\",\"transferPhoneNumber\":\"\",\"transferReason\":\"\"," +
		"\"callDuration\":82,\"customerId\":\"febd313261f42e5a35581faf5270165\",\"customerName\":\"时科-长弓-58同城-F1\"," +
		"\"openId\":\"54c82a5dd7ab4a9ab9ec3d10c32bef11\",\"smsContent\":null,\"ability\":\"10\",\"smsCount\":null," +
		"\"charge6\":14,\"charge60\":2} "
	var bill_ware_house Hd
	err = json.Unmarshal([]byte(jsonString), &bill_ware_house)
	if err != nil {
		log.Fatalln(err)
		return
	}

	sqlStr := "insert into bill_ware_house (`service_code`,`phone_number_a`,`phone_number_a_area_code`," +
		"`phone_number_a_operator`,`phone_number_a_province`,`phone_number_a_city`,`phone_number_b`," +
		"`phone_number_b_area_code`,`phone_number_b_operator`,`phone_number_b_province`,`phone_number_b_city`," +
		"`phone_number_x`,`phone_number_x_area_code`,`phone_number_x_operator`,`phone_number_x_province`," +
		"`phone_number_x_city`,`phone_number_y`,`phone_number_y_area_code`,`phone_number_y_operator`," +
		"`phone_number_y_province`,`phone_number_y_city`,`extension_number`,`binding_id`,`call_id`,`call_time`," +
		"`ringing_time`,`start_time`,`release_time`,`release_direction`,`release_cause`,`call_recording`," +
		"`recording_url`,`recording_mode`,`call_type`,`call_result`,`transfer_phone_number`,`transfer_reason`," +
		"`call_duration`,`customer_id`,`customer_name`,`open_id`,`sms_content`,`ability`,`sms_count`,`charge6`," +
		"`charge60`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(bill_ware_house.ServiceCode, bill_ware_house.PhoneNumberA, bill_ware_house.PhoneNumberAAreaCode, ""+
		bill_ware_house.PhoneNumberAOperator, bill_ware_house.PhoneNumberAProvince, bill_ware_house.PhoneNumberACity, ""+
		bill_ware_house.PhoneNumberB, bill_ware_house.PhoneNumberBAreaCode, bill_ware_house.PhoneNumberBOperator, ""+
		bill_ware_house.PhoneNumberBProvince, bill_ware_house.PhoneNumberBCity, bill_ware_house.PhoneNumberX, ""+
		bill_ware_house.PhoneNumberXAreaCode, bill_ware_house.PhoneNumberXOperator, bill_ware_house.PhoneNumberXProvince, ""+
		bill_ware_house.PhoneNumberXCity, bill_ware_house.PhoneNumberY, bill_ware_house.PhoneNumberYAreaCode, ""+
		bill_ware_house.PhoneNumberYOperator, bill_ware_house.PhoneNumberYProvince, bill_ware_house.PhoneNumberYCity, ""+
		bill_ware_house.ExtensionNumber, bill_ware_house.BindingId, bill_ware_house.CallId, bill_ware_house.CallTime, ""+
		bill_ware_house.RingingTime, bill_ware_house.StartTime, bill_ware_house.ReleaseTime, bill_ware_house.ReleaseDirection, ""+
		bill_ware_house.ReleaseCause, bill_ware_house.CallRecording, bill_ware_house.RecordingUrl, bill_ware_house.RecordingMode, ""+
		bill_ware_house.CallType, bill_ware_house.CallResult, bill_ware_house.TransferPhoneNumber, ""+
		bill_ware_house.TransferReason, bill_ware_house.CallDuration, bill_ware_house.CustomerId, ""+
		bill_ware_house.CustomerName, bill_ware_house.OpenId, bill_ware_house.SmsContent, ""+
		bill_ware_house.Ability, bill_ware_house.SmsCount, bill_ware_house.Charge6, bill_ware_house.Charge60)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println("insert success!")
}
