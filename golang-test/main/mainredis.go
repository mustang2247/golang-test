package main

import (
	"fmt"
	"github.com/hoisie/redis"
	"sync"
)

func main() {
	initRedis()
	// 测试string
	client.Set("s", []byte("hello world"))
	val, _ := client.Get("s")
	fmt.Println(string(val))
	client.Del("s")

	//测试list
	vals := []string{"a1", "b2", "c3", "d4", "e5"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("l")
}

var (
	client *redis.Client
	mutex  sync.Mutex
)

func initRedis() {
	mutex.Lock()
	defer mutex.Unlock()
	if client != nil {
		return
	}
	client = &redis.Client{
		Addr:        "127.0.0.1:6379",
		Db:          0,
		Password:    "",
		MaxPoolSize: 10000,
	}
	if err := client.Auth(""); err != nil {
		fmt.Println("Auth:", err.Error())
		return
	}
}