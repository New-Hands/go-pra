package pra_io

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"testing"
)

// 标准输入 输出
// https://flaviocopes.com/go-tutorial-lolcat/
func TestInOut(t *testing.T) {
	term.NewTerminal()
	stdin := os.Stdin
	read, err := stdin.Read(make([]byte, 1, 1))
	if err != nil {
		return
	}
	fmt.Println(read)
	reader := bufio.NewReader(stdin)
	readString, err := reader.ReadString('d')
	if err != nil {
		return
	}

	fmt.Println(readString)
}
