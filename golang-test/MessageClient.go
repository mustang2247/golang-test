package main

import (
	"net"
	"fmt"
	"utils/logutil"
	"message"
)

func main() {
	go message.StartServer(":1171", "tcp")

	fmt.Println("begin conn")
	tcpAddress, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1171")
	logutil.CheckErr(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	logutil.CheckErr(err)

	writeChan := make(chan []byte, 1024)
	readChan := make(chan []byte, 1024)

	go writeConnection(conn, writeChan)
	go readConnection(conn, readChan)

	//go handleReadChannel(readChan)

	for {
		var s string
		fmt.Scan(&s)
		writeChan <- []byte(s)
	}

}

func readConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		logutil.CheckErr(err)
		println("Received from:", conn.RemoteAddr(), string(buffer[:n]))
		//channel <- buffer[:n]
	}

}

func writeConnection(conn *net.TCPConn, channel chan []byte) {
	defer conn.Close()
	for {
		select {
		case data := <- channel:
			_, err := conn.Write(data)
			logutil.CheckErr(err)
			println("Write to:", conn.RemoteAddr(), string(data))
		}
	}
}