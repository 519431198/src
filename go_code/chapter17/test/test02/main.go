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
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis-set err=", err)
		return
	}
}
