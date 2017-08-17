// jsonparse project main.go
package main

import (
	"utils/jsonhelper"
	"fmt"
)

type Server struct {
	serverName string `json:"serverName"`
	ServerIP string `json:"serverIP"`
}

type ServersSlice struct {
	Servers []Server `json:"servers"`
	Desc string   `json:"desc"`
}

func main() {
	servers := []Server{{serverName: "admin", ServerIP: "192.168.20.131"}, {serverName: "admin1", ServerIP: "192.168.20.132"}}
	desc := "描述部分"
	var ss ServersSlice
	ss.Desc = desc
	ss.Servers = servers
	jsonhelper.Marshal("server.json", ss)

	//=========================
	s := ServersSlice{}
	jsonhelper.Open("server.json", &s)
	fmt.Println("\n")
	fmt.Printf("%#v", s)


}
