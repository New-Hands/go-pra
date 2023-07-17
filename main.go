package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString, err1 := reader.ReadString('d')
	if err1 != nil {
		return
	}

	fmt.Println(readString)
}
