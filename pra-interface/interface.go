package pra_interface

import "fmt"

type it2 interface {
	Hello()
}

type it1 interface {
	Hello()
}

type sToIt struct {
}

// Hello 没有明确的实现关系指明
func (sToIt) Hello() {
	fmt.Println("hello world")
}
