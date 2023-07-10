package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	md52 "crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 客户openId
const openId = "030f2505c9aa4ceb85afb8291e60eea9"

// 客户密匙
const customerSecret = "685a9cefb756b3128235d7bd02f7ebeb"

// 功能编码
const functionId = "AX_BINDING"

// 是否加密
const encrypt = false

// 请求参数
var params = map[string]interface{}{
	"phoneNumberX":     "15648117254",
	"phoneNumberA":     "17704112288",
	"areaCode":         "0471",
	"name":             "赵四",
	"cardType":         "身份证",
	"cardNumber":       "身份证号码",
	"expiration":       26000,
	"callRecording":    "1",
	"callDisplay":      "1",
	"smsMtChannel":     "3",
	"calledAudioCode":  "1",
	"callingAudioCode": "",
}

// md5计算
func md5(str string) string {
	h := md52.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// tripleDES加密
func tripleDES(str, key string) (string, error) {
	data := []byte(str)
	keyBytes := []byte(key[8:24] + "00000000")
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		return "", err
	}
	padSize := des.BlockSize - (len(data) % des.BlockSize)
	padding := make([]byte, padSize)
	for i := range padding {
		padding[i] = byte(padSize)
	}
	data = append(data, padding...)
	encrypted := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, keyBytes[:des.BlockSize])
	mode.CryptBlocks(encrypted, data)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 签名
func sign(param map[string]interface{}, secret string) string {
	signature := param["openId"].(string) +
		fmt.Sprintf("%v", param["timeStamp"]) +
		fmt.Sprintf("%v", param["id"]) +
		md5(secret)

	requestUnits := param["request"].([]interface{})
	for _, unit := range requestUnits {
		unitMap := unit.(map[string]interface{})
		signature += unitMap["id"].(string)
		signature += unitMap["data"].(string)
	}
	return md5(signature)
}

// base64 加密
func base64Encode(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

func main() {
	// 定义时间戳
	timeStamp := time.Now().Unix()

	dataBytes, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error marshaling params:", err)
		return
	}
	data := string(dataBytes)

	if encrypt {
		encData, err := tripleDES(data, md5(customerSecret))
		if err != nil {
			fmt.Println("Error encrypting data:", err)
			return
		}
		data = base64Encode(encData)
	}

	// 定义请求
	request := map[string]interface{}{
		"id":        timeStamp,
		"openId":    openId,
		"request":   []interface{}{map[string]interface{}{"id": functionId, "encrypt": encrypt, "data": data}},
		"timeStamp": timeStamp,
	}
	request["signature"] = sign(request, customerSecret)

	requestStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return
	}

	SendRequest(requestStr, "http://61.139.144.36:9000/bas/api/query/request.json")
	//fmt.Println(string(requestStr))
}

func SendRequest(requestStr []byte, url string) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
