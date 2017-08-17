package main

import (
	"fmt"
	"time"
)

func main()  {
	// go routine
	go Go()
	time.Sleep(2 * time.Second)
	fmt.Println("hell")
}

func Go() {
	fmt.Println("gogo go")
}