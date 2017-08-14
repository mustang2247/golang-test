package main

import (

	"fmt"
	"db/redis"
)


func main() {
	test(1, 23, 345)
}

func test(args ...interface{})  {
	for key, value := range args {
		fmt.Println(key, value)
	}

	//client = redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})

	redis.Open("127.0.0.1:6379", "", 0)
	redis.Set("testgo", "gotest 哈哈哈哈哈哈哈")
	res,err := redis.Get("testgo")
	if err != nil {
		println("mysql checkErr: " + err.Error())
		panic(err)
	}

	println(res)

}

