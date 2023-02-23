package main

import (
	"encoding/json"
	"fmt"
)

type log struct {
	ServiceCode          string
	PhoneNumberA         string
	PhoneNumberAAreaCode string
	PhoneNumberAOperator string
}

func testStruct() {
	log := log{
		ServiceCode:          "2",
		PhoneNumberA:         "02569845957",
		PhoneNumberAAreaCode: "25",
		PhoneNumberAOperator: "固话",
	}
	data, err := json.Marshal(&log)
	if err != nil {
		fmt.Println("marshal err=", err)
	}
	fmt.Printf("%v", string(data))
}

func main() {
	testStruct()
}
