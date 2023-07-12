package cmd

import (
	"fmt"
	"net"
	"testing"
)

func TestListen(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:6001")
	if err != nil {
		return
	}
	for true {
		fmt.Println("new to accept")
		client, err := listen.Accept()
		fmt.Println("accept client", client.RemoteAddr())
		if err != nil {
			return
		}
		write, wErr := client.Write([]byte("hello"))
		if wErr != nil {
			return
		}

		_ = fmt.Sprintf("write to %d bytes", write)
		_ = client.Close()
	}
}
