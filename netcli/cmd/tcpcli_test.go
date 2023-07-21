package cmd

import (
	"bytes"
	"fmt"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	dial, err := net.Dial("tcp", "10.194.50.30:6112")
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
			Ip:             "10.194.50.30",
			Port:           6112,
			ReceiveTimeOut: 10,
		},
	}

	_ = tcp.Start()

	go func() {
		for true {
			time.Sleep(time.Millisecond * 10)
			_ = tcp.Write(&model.MsgTo{
				Data: []byte("hello" + strconv.Itoa(time.Now().Second())),
			})
		}

	}()
	time.Sleep(time.Minute * 10)
}

func TestWrapClient2(t *testing.T) {
	tcp := &Tcp{
		NetParam: model.NetParam{
			Ip:             "10.194.50.30",
			Port:           6112,
			ReceiveTimeOut: 10,
		},
	}

	_ = tcp.Start()

	go func() {
		for true {
			time.Sleep(time.Millisecond * 10)
			_ = tcp.Write(&model.MsgTo{
				Data: []byte("hello" + strconv.Itoa(time.Now().Second())),
			})
		}

	}()
	time.Sleep(time.Minute * 10)
}
