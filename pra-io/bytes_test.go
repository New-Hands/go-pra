package pra_io

import (
	"bytes"
	"fmt"
	"testing"
)

// 关于io接口 定义在 io包中
// 包括
// type Reader interface

// 字节变形操作
func TestBytesOp(t *testing.T) {
	// len cap
	i := make([]byte, 30, 30)
	fmt.Println(len(i), cap(i))
	fmt.Println(i)

	// slice op 的内存共享 超过cap容量会重新分配内存
	// s := arr[startIndex:endIndex]
	// s := arr[startIndex:]
	// s := arr[:endIndex]
	i2 := i[:20]
	i[19] = 19
	fmt.Println(i2)
	fmt.Println(i)

	// append 超出容量导致重新分配内存
	i3 := append(i, 1)
	i[19] = 20
	// i3 不受 i影响
	fmt.Println(i)
	fmt.Println(i3)

	// FULL LIMITED 限制内存容量切片 切片重新分配容量
	//i4 := i[:10:10]

}

// read to buff
func TestReadToBuff(t *testing.T) {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")

	bytes.NewReader([]byte(""))

	fmt.Println(buf.Bytes())
}

// write to bytes
func TestWriteToBytes(t *testing.T) {
}
