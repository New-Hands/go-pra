package cmd

import (
	"bytes"
	"fmt"
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
