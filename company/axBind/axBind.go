package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/my/repo/company/axBind/utils"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 定义常量
const url = "https://shoujihao.uutool.cn/no/"

// 客户信息
const (
	openId         = "030f2505c9aa4ceb85afb8291e60eea9"
	customerSecret = "685a9cefb756b3128235d7bd02f7ebeb"
	functionId     = "AX_BINDING"
	encrypt        = false
)

// 请求参数
var params = map[string]interface{}{
	"phoneNumberX":     "",
	"phoneNumberA":     "",
	"areaCode":         "",
	"name":             "赵宇鋆",
	"cardType":         "身份证",
	"cardNumber":       "330126198612053338",
	"callDisplay":      "1",
	"expiration":       10,
	"callRecording":    "1",
	"smsMtChannel":     "3",
	"calledAudioCode":  "",
	"callingAudioCode": "",
}

// ApiTools 包含用于签名和加密的函数
type ApiTools struct{}

func main() {
	apiTools := &ApiTools{}
	//phone := map[string]string{
	//	"15555488644": "05971140201",
	//	"13001140205": "05971140205",
	//	"13001140206": "05971140206",
	//	"13001140208": "05971140208",
	//	"13001140211": "05971140211",
	//	"13001140212": "05971140212",
	//}
	//
	//for numX, numA := range phone {
	//
	//	webSiteData := utils.GetWebsiteData(url + numX)
	//	areaCode := utils.Code(string(webSiteData))
	//	params["phoneNumberA"] = numX
	//	params["phoneNumberX"] = numA
	//	params["areaCode"] = areaCode
	//
	//	timeStamp, data := apiTools.TimeStampMd5Data()
	//	// 定义请求
	//	requestStr := apiTools.ReqStr(timeStamp, data)
	//	fmt.Println(string(requestStr))
	//	// 发送请求
	//	//apiTools.Req(requestStr)
	//
	//	time.Sleep(time.Second * 2)
	//}

	// 读文件方式绑定
	phone := map[string]string{}
	data, _ := os.ReadFile("test/test03/phone.yaml")
	_ = yaml.Unmarshal(data, &phone)
	for numX, numA := range phone {
		webSiteData := utils.GetWebsiteData(url + numX)
		areaCode := utils.Code(string(webSiteData))
		params["phoneNumberA"] = numX
		params["phoneNumberX"] = numA
		params["areaCode"] = areaCode

		timeStamp, data := apiTools.TimeStampMd5Data()
		// 定义请求
		requestStr := apiTools.ReqStr(timeStamp, data)
		fmt.Println(string(requestStr))
		// 发送请求
		//apiTools.Req(requestStr)

		time.Sleep(time.Second * 2)
	}
}

// TimeStampMd5Data 定义时间戳及加密
func (t *ApiTools) TimeStampMd5Data() (string, string) {
	// 定义时间戳
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	dataBytes, _ := json.Marshal(params)
	data := string(dataBytes)
	if encrypt {
		key := []byte(t.MD5(customerSecret)[8:24] + "00000000")
		data, _ = t.TripleDES(data, key)
	}
	return timeStamp, data
}

// ReqStr 定义请求
func (t *ApiTools) ReqStr(timeStamp, data string) []byte {
	request := map[string]interface{}{
		"id":     timeStamp,
		"openId": openId,
		"request": []map[string]interface{}{
			{
				"id":      functionId,
				"encrypt": encrypt,
				"data":    data,
			},
		},
		"timeStamp": timeStamp,
	}
	request["signature"] = t.Sign(request)

	reqStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println(errors.New("请求序列化失败"))
	}
	return reqStr
}

// Req 发送请求
func (t *ApiTools) Req(reqStr []byte) {
	// 发送请求
	req, _ := http.NewRequest("POST", "http://61.139.144.36:9000/bas/api/query/request.json", bytes.NewBuffer(reqStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// 读取响应
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// MD5 计算字符串的 MD5 值
func (t *ApiTools) MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// TripleDES 使用 Triple DES 算法加密字符串
func (t *ApiTools) TripleDES(str string, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	// PKCS7Padding
	padding := des.BlockSize - len(str)%des.BlockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	str += string(padtext)

	// 加密
	ciphertext := make([]byte, len(str))
	mode := cipher.NewCBCEncrypter(block, key[:des.BlockSize])
	mode.CryptBlocks(ciphertext, []byte(str))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Sign 计算请求的签名
func (t *ApiTools) Sign(param map[string]interface{}) string {
	signature := param["openId"].(string) +
		param["timeStamp"].(string) +
		param["id"].(string) +
		t.MD5(customerSecret)

	requestUnits := param["request"].([]map[string]interface{})
	for _, unit := range requestUnits {
		signature += unit["id"].(string)
		signature += unit["data"].(string)
	}

	return t.MD5(signature)
}
