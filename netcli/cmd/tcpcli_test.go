package cmd

import (
	"bytes"
	"fmt"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	dial, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(make([]byte, 10))

	read, err := dial.Read(buffer.Bytes())
	if err != nil {
		return
	}

	readBytes, rErr := buffer.ReadBytes(byte(read))
	if nil != rErr {
		fmt.Println(rErr)
	}

	fmt.Println(string(readBytes[:read]))
}

func TestWrapClient(t *testing.T) {
	tcp := &Tcp{
		NetParam: model.NetParam{
			Ip:             "127.0.0.1",
			Port:           6001,
			ReceiveTimeOut: 10,
		},
	}

	err := tcp.Start()
	if err != nil {
		return
	}
	read, err := tcp.Read()
	if err != nil {
		return
	}
	fmt.Println(read)
}
