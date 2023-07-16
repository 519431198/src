package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func main() {
	phone := map[string]string{}
	data, _ := os.ReadFile("test/test03/phone.yaml")
	_ = yaml.Unmarshal(data, &phone)
	for i, v := range phone {
		fmt.Printf("phoneNumX: %s,phoneNmuA: %s\n", i, v)
		time.Sleep(time.Second * 2)
	}
}
