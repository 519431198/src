package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Db, err := sql.Open("mysql", "root:123@tcp(10.0.0.10:3306)/bill")
	if err != nil {
		panic(err.Error())
	}
	// 1.写 sql 语句
	sqlStr := "SELECT sum(charge6),sum(charge60) FROM `bill_ware_house_2023-06-18` WHERE phone_number_x IN ('18613043516') AND charge60 != '0'"
	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var charge6 int
		var charge60 int
		rows.Scan(&charge6, &charge60)
		fmt.Println(charge6, charge60)
	}

}
