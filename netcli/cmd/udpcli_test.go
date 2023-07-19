package cmd

import (
	"fmt"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"testing"
)

// go udp
// udp 无连接 不区分服务端客户端 但是go中 有两种不同的UdpConn构造模式 对应服务端模式  客户端模式
// 客户端创建方式构建的UdpConn 虽然也能指定监听端口，但是对监听连接做了限制 指定了外部地址为构建客户端的远端地址
// 服务端模式能实现 接收数据 发送数据功能，类型原生
func TestUdp(t *testing.T) {
	udp, err := net.DialUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("*"),
		Port: 6114,
	}, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 666,
	})
	if err != nil {
		return
	}

	listen()

	fromUDP, u, rErr := udp.ReadFromUDP(make([]byte, 2))
	if err != nil {
		fmt.Println(rErr)
	}

	fmt.Println(fromUDP, u)
}

func listen() {
	listenUDP, err1 := net.ListenUDP("udp", &net.UDPAddr{
		Port: 6112,
	})
	if err1 != nil {
		fmt.Println(err1)
	}

	bytes := make([]byte, 65535, 65535)

	for true {
		readFromUDP, u2, err2 := listenUDP.ReadFromUDP(bytes)
		fmt.Printf("%p", bytes)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(u2, string(bytes[:readFromUDP]))
		// echo
		_, wErr := listenUDP.WriteToUDP([]byte("hello"), u2)
		if wErr != nil {
			fmt.Println(wErr)
		}
	}

}

func TestUdpCli(t *testing.T) {
	// send to 127.0.0.1 6112 and bind
	udp := Udp{
		NetParam: model.NetParam{
			Ip:         "127.0.0.1",
			Port:       6114,
			ListenPort: 666,
		},
	}

	err := udp.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(udp.Read())
}
