package main

import "github.com/go-redis/redis"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "10.0.0.4:6379",
		DB:   0,
	})
	defer rdb.Close()
	//rdb.Do("set", "hug", "yici")
	rdb.LPush("nihao", "4")
}
