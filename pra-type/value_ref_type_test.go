package pra_type

import (
	"fmt"
	"testing"
)

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

func TestStruct(t *testing.T) {
	tt := Ts{n: "type ts"}

	fmt.Printf("%p \n", tt)
}

func TestSlice(t *testing.T) {
	bytes := [...]byte{'a', 'b', 'c'}
	fmt.Printf("%p \n", bytes)
}
