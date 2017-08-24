// go 并发
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i<10; i++ {
		wg := sync.WaitGroup{}
		wg.Add(10)
		for i := 0; i<10; i++ {
			go Go(&wg, i)
		}
		wg.Wait()

		fmt.Println(" \n ")
	}
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}