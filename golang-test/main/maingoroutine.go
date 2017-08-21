// go 并发
package main

import (
	"fmt"
)

func main()  {
	// go routine
	//go Go()
	//time.Sleep(2 * time.Second)
	//fmt.Println("hell")

	c := make(chan bool)
	go func() {
		fmt.Println("gogogo")
		c <- true
		close(c)//使用for range必须明确关闭
	}()
	//<-c

	for v := range c {
		fmt.Println(v)
	}

	// 单项通道和双向通道
}

func Go() {
	fmt.Println("gogo go")
}