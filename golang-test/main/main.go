package main

import (

	"fmt"
	"github.com/mustang2247/golang-utils/db/redis"
)


func main() {
	test(1, 23, 345)
}

func test(args ...interface{})  {
	for key, value := range args {
		fmt.Println(key, value)
	}

	redis.Open("127.0.0.1:6379", "", 0)
	redis.Set("testgo", "gotest 哈哈哈哈哈哈哈")
	res,err := redis.Get("testgo")
	if err != nil {
		println("mysql checkErr: " + err.Error())
		panic(err)
	}

	println(res)

}

