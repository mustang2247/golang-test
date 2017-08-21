package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"
	"strconv"
)

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//goroutine 实现多并发
		go handleClient(conn)
	}
}

// 实现多并发执行
func handleClient(conn net.Conn) {
	//defer conn.Close()
	//daytime := time.Now().String()
	//conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client

	// conn.SetReadDeadline()设置超时 2分钟
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	// 需要指定一个最大长度以防止flood attack
	request := make([]byte, 128) // set maxium request length to 128B to prevent flood attack
	defer conn.Close()  // close connection before exit
	for {
		// 使用conn.Read()不断读取客户端发来的请求
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			fmt.Println(string(request[:read_len]))
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		// 每次读取到请求处理完毕后，需要清理request
		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
