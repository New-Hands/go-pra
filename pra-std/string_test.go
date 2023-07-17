package pra_std

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	// 字符串拼接
	var str string = "ddd"
	fmt.Printf("%d \n", unsafe.Sizeof(str))
	str = "dddddddd"
	fmt.Printf("%d \n", unsafe.Sizeof(str))

}
