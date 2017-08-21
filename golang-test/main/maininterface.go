package main

import "fmt"

//==========interface=============

type empty interface {

}

type USB interface {
	Name() string
	Use() string
	// 嵌入接口
	Connecter
}

type Connecter interface {
	Connect()
}

//===========接口实现============

type PhoneConnecter struct {
	name string
	use string
}

func (pc PhoneConnecter) Name() string  {
	return pc.name
}

func (pc PhoneConnecter) Use() string  {
	return pc.use
}

func (pc PhoneConnecter) Connect()  {
	fmt.Println("Connect: ", pc.name, pc.use)
}

type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("Conncet: ", tv.name)
}


func main() {
	pc := PhoneConnecter{"PhoneConnecter", "Connect"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	Disconnect(a)

	tv := TVConnecter{"TVConnecter"}
	tv.Connect()
	Disconnect(tv)
}

// 类型断言
// 空接口
func Disconnect(usb interface{}) {
	//if pc, ok := usb.(PhoneConnecter); ok {
	//	fmt.Println("Disconnect: ", pc.name)
	//	return
	//}

	switch v:=usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect: ", v.name)
	default:
		fmt.Println("Unknown device.")
	}


}