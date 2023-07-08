package main

import (
	"fmt"
	"github.com/tjfoc/gmtls"
)

func main() {
	var l, err = gmtls.Listen("tcp", "127.0.0.1", nil)
	if nil != err {
		fmt.Printf(err.Error())
	}

	l.Accept()
}
