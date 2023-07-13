package pra_mem

import (
	"fmt"
	"testing"
)

type MemTest struct {
	str     string
	intT    int
	obj     map[string]string
	intAddr *int
}

// 变量及数据本身？

// 指针是一个基础类型，能被copy
// 引用类型 值是地址 不会对地址值进行拷贝
// 测试传参类型 传值 传引用？ 都是传值 都会拷贝变量
func TestParamCopy(t *testing.T) {
	obj := make(map[string]string, 10)
	obj["tt"] = "tt"

	i := 8

	te := MemTest{
		str:     "str",
		intT:    1,
		obj:     obj,
		intAddr: &i,
	}
	// {str 1 map[tt:tt-copy]} string 也是基础类型？
	// map值被修改
	funCall(te)

	// 被拷贝 不一样的地址 值类型数据 引用类型数据（数据是数据的地址）
	// 值数据传递会copy值  地址数据传递copy地址值 所以也是同一份数据
	//str 0xc000132900
	//obj 0xc0001328d0
	//intAddr 0xc00011a2c0
	//str new 0xc000132930

	// 注意 区别变量地址 及  地址数据（数据是地址值）
	//obj new 0xc0001328d0
	//intAddr new 0xc00011a2c0
	fmt.Printf("str %p\n", &te.str)
	fmt.Printf("obj %p\n", te.obj)
	fmt.Printf("intAddr %p\n", te.intAddr)
	fmt.Printf("intAddr str %p\n", &te.intAddr)
	funCall2(te)
}

func funCall(test MemTest) {
	// 修改值
	test.obj["tt"] = "tt-copy"
	test.str = "srt-copy"
	test.intT = 2
}

func funCall2(test MemTest) {
	// 修改值
	fmt.Printf("str new %p\n", &test.str)
	fmt.Printf("obj new %p\n", test.obj)
	fmt.Printf("intAddr new %p\n", test.intAddr)
	fmt.Printf("intAddr str new %p\n", &test.intAddr)
}
