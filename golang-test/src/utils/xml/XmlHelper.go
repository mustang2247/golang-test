package xmlhelper

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)


type Recurlyservers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct{
	XMLName xml.Name `xml:"server"`
	Desc string `xml:"desc,attr"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
	TestDesc string `xml:",innerxml"`
}

func Open(path string, value interface{})  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	err = xml.Unmarshal(data, value)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	fmt.Printf("%#v", value)
}