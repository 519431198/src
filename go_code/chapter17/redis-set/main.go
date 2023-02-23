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
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis-set err=", err)
		return
	}
	err = rdb.Set("name", "tomejerry", 0).Err()
	if err != nil {
		fmt.Println("rdb.ser err =", err)
	}
	val, err := rdb.Get("name").Result()
	fmt.Println(val)
	if err != nil {
		fmt.Println("Get err=", err)
	}
}
