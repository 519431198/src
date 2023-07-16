package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// 客户信息
const (
	openId         = "dc0e489d826c4954a81983fd9ada887d"
	customerSecret = "3acf66cf474a400bb746bf91c5209a1d"
	functionId     = "AXB_BINDING"
	encrypt        = false
)

// 请求参数
var params = map[string]interface{}{
	"phoneNumberA":  "05973377454",
	"phoneNumberX":  "17602888546",
	"phoneNumberB":  "18930547397",
	"audioCode":     "0,0,0",
	"areaCode":      "028",
	"expiration":    10,
	"callRecording": "0",
	"callDisplay":   "0,1",
	"callRestrict":  "1",
	"smsMtChannel":  "3",
	"smsSuffix":     "",
}

// ApiTools 包含用于签名和加密的函数
type ApiTools struct{}

func main() {
	apiTools := ApiTools{}

	// 定义时间戳
	// 定义时间戳
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	dataBytes, _ := json.Marshal(params)
	data := string(dataBytes)
	if encrypt {
		key := []byte(apiTools.MD5(customerSecret)[8:24] + "00000000")
		data, _ = apiTools.TripleDES(data, key)
	}

	// 定义请求
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
	request["signature"] = apiTools.Sign(request)

	requestStr, _ := json.Marshal(request)
	//fmt.Println(string(requestStr))
	apiTools.Req(requestStr)
}

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
