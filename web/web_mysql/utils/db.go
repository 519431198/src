package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123@tcp(10.0.0.10:3306)/bill")
	if err != nil {
		panic(err.Error())
	}
}
