package pra_type

import (
	"fmt"
	"testing"
)

func TestValueAndRef(t *testing.T) {
	// Strings are immutable, so you cannot update the value of a string.
	//
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
