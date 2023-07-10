package pra_interface

import (
	"fmt"
	"reflect"
	"testing"
)

// interface 方法集
// 如果两个接口有一个相同方法 结构体是继承了哪个interface
// 没有明确的实现关系指明 自定义type 可以被赋值给 所有 自己实现过方法的 interface value
func TestItFun(t *testing.T) {
	s := sToIt{}
	var it1 it1 = s
	it1.Hello()
	var it2 it2 = s
	it2.Hello()

	// nil interface value
	//it1 = nil
	//it1.Hello()

	// 空  interface{}
	// An empty interface may hold values of any type
	var eiter interface{}
	// output nil
	fmt.Println(eiter)
	eiter = 42
	fmt.Println(eiter)
	eiter = "ddd"
	fmt.Println(eiter)
	eiter = s
	fmt.Println(eiter)

	// 类型断言
	s2, ok := eiter.(string)
	fmt.Println(s2, ok)
	// panic
	//i2 := eiter.(string)

	switch eiter.(type) {
	case string:
		fmt.Println()
	case int:
		fmt.Println("int")
	default:
	}

	// 使用反射获取接口类型
	var a interface{} = "Hello, world!"
	tt := reflect.TypeOf(a)
	fmt.Println(tt)
}

// 熟悉有哪些 内置接口
