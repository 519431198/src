package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "10.0.0.4:6379",
		DB:   0,
	})
	defer rdb.Close()
	rdb.HSet("user", "name", "tom")
	val, err := rdb.HGet("user", "name").Result()
	if err != nil {
		fmt.Println("rdb.hget err=", err)
	}
	fmt.Println(val)

	var adr = make(map[string]interface{})
	adr["name1"] = "beijing"
	adr["name2"] = "tianjing"
	res, err := rdb.HMSet("user02", adr).Result()
	fmt.Println(res)
	val1, err := rdb.HMGet("user02", "name1", "name2").Result()
	for i, v := range val1 {
		fmt.Printf("r[%d]=%s", i, v)
	}
}
