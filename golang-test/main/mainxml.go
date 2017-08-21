package main

import (
	"encoding/xml"
	"github.com/mustang2247/golang-utils/utils/xmlhelper"
	"fmt"
)

//type Recurlyservers struct{
//	XMLName xml.Name `xml:"servers"`
//	Version string `xml:"version,attr"`
//	Svs []server `xml:"server"`
//	Description string `xml:",innerxml"`
//}
//
//type server struct{
//	XMLName xml.Name `xml:"server"`
//	Desc string `xml:"desc,attr"`
//	ServerName string `xml:"serverName"`
//	ServerIP string `xml:"serverIP"`
//	TestDesc string `xml:",innerxml"`
//}


type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}

type server struct{
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	//v := Recurlyservers{}
	//xmlhelper.Open("servers.xml", &v)
	////fmt.Printf("%#v", v)
	//fmt.Println(v.Version + "\n")
	//for _, val := range v.Svs{
	//	fmt.Println(val.ServerName, val.ServerIP)
	//}

	v := &Servers{Version: "1.0"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.1"})
	xmlhelper.MakeXml("servers.xml", &v)
	fmt.Println(v)
}
