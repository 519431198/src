package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	req := Webhook{
		AccessToken: "f9f7e29c56a678facf1cafcc7bdd5263a4bef01e4335dc97436cb131c33bd439",
		Secret:      "SEC6d80d653ee3f110571b91c55320cb57f2ee89f9471317c46a6689570adb13c2a",
		//艾特所有人需要两个同时开启
		//EnableAt: true, // 开启艾特
		//AtAll:       true, // 艾特所有人
	}
	_ = req.SendMessage("今天天气真好!", "18326147303")
}

type Webhook struct {
	AccessToken string
	Secret      string
	EnableAt    bool
	AtAll       bool
}

// SendMessage Function to send message
//goland:noinspection GoUnhandledErrorResult
func (t *Webhook) SendMessage(s string, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": s,
		},
	}
	if t.EnableAt {
		if t.AtAll {
			if len(at) > 0 {
				return errors.New("the parameter \"AtAll\" is \"true\", but the \"at\" parameter of SendMessage is not empty")
			}
			msg["at"] = map[string]interface{}{
				"isAtAll": t.AtAll,
			}
		} else {
			msg["at"] = map[string]interface{}{
				"atMobiles": at,
				"isAtAll":   t.AtAll,
			}
		}
	} else {
		if len(at) > 0 {
			return errors.New("the parameter \"EnableAt\" is \"false\", but the \"at\" parameter of SendMessage is not empty")
		}
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (t *Webhook) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *Webhook) getURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + t.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, t.Secret)
	sign := t.hmacSha256(stringToSign, t.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
