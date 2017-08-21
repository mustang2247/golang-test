package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError2(err)

	conn, err := net.DialTCP("tcp", nil, udpAddr)
	checkError2(err)

	_, err = conn.Write([]byte("anything"))
	checkError2(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError2(err)

	fmt.Println(string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}

func checkError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}