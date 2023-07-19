package pra_type

import (
	"fmt"
	"testing"
	"unsafe"
)

// // Assume T is an arbitrary type and Tkey is
//// a type supporting comparison (== and !=).
//
//*T         // a pointer type
//[5]T       // an array type
//[]T        // a slice type
//map[Tkey]T // a map type
//
//// a struct type
//struct {
//	name string
//	age  int
//}
//
//// a function type
//func(int) (bool, string)
//
//// an interface type
//interface {
//	Method0(string) int
//	Method1() (int, bool)
//}
//
//// some channel types
//chan T
//chan<- T
//<-chan T

//  the difference between Unicode and UTF-8,
// the difference between a string and a string literal, and other even more subtle distinctions.

func TestValueAndRef(t *testing.T) {
	// Strings are immutable, so you cannot update the value of a string.
	// a string is in effect a read-only slice of bytes.
	var first = "test"
	var second = first
	first = "another test"
	//	first  //"another test"
	//	second //"test"
	m := map[string]string{}
	fmt.Printf("%p \n", m)
	fmt.Printf("%p %s \n", &first, first)
	fmt.Printf("%p %s \n", &second, second)
}

func TestStringArr(t *testing.T) {
}

type S1 struct {
	i1  int32
	str []string
}
type S2 struct {
	i1  int64
	str []string
}

func TestStruct(t *testing.T) {
	// 内存对齐
	fmt.Println(unsafe.Sizeof(S1{}))
	fmt.Println(unsafe.Sizeof(S2{}))

	// pointer
	fmt.Println(unsafe.Sizeof(&S1{}))
	fmt.Println(unsafe.Sizeof(&S2{}))
	m := map[string]int{
		"1": 1,
		"3": 1,
		"2": 1,
	}
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(&m))

	// slice
	fmt.Println("slice size")
	fmt.Println(unsafe.Sizeof([5]string{}))
	fmt.Println(unsafe.Sizeof([2]string{"dd", "dd"}))

	// array
	fmt.Println("array size")
	fmt.Println(unsafe.Sizeof([]string{}))
	fmt.Println(unsafe.Sizeof([]string{"dd", "dd"}))

	fmt.Println("string var size", unsafe.Sizeof("ddd"))
	fmt.Println("string var size", unsafe.Sizeof("dddd"))

	fmt.Println(unsafe.Sizeof([]int{}))
	fmt.Println(unsafe.Sizeof('b'))

}

func TestString(t *testing.T) {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")
	fmt.Println("Printf with %x:")
	fmt.Printf("%x \n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
	//UTF-8 and string literals
	fmt.Printf("%s", "ddd")
}

type Ts struct {
	n string
}

func TestSlice(t *testing.T) {
	bytes := [...]byte{'a', 'b', 'c'}
	fmt.Printf("%p \n", bytes)
}
