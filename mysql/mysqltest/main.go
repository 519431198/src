package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wanghuiyt/ding"
	"log"
	"strconv"
)

type Hd struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Charge6      int    `json:"charge6"`
	Charge60     int    `json:"charge60"`
	CallTotal    int    `json:"call_total"`
}

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10:3306)/bill")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from bill_ware_house_count where customer_name like '时科-%' order by	Charge60 desc ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var message = make([]string, 20, 20)
	for rows.Next() {
		hd := new(Hd)
		var err = rows.Scan(&hd.Id, &hd.CustomerName, &hd.Charge6, &hd.Charge60, &hd.CallTotal)
		if err != nil {
			panic(err.Error())
		}
		var str string
		str = fmt.Sprintf("%s\t%s\n", strconv.Itoa(hd.Charge60), hd.CustomerName)
		message = append(message, str)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	var str string
	for _, v := range message {
		str += v
	}

	d := ding.Webhook{
		AccessToken: "0b8e5c9299cc7ca77cc863381e8b5949c648a4804ae3993b046e3ee4ab33b70d",
		Secret:      "SEC4b3633a968fe18595a855597c2be8315917e26b08b2eafe2f82f4d40314271eb",
	}
	err = d.SendMessageText(str, "*")
	if err != nil {
		log.Fatalln(err)
	}
}
