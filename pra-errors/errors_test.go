package pra_errors

import (
	"fmt"
	"testing"
)

// error is a builtin
func TestError(t *testing.T) {
	fmt.Println(ReError())
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return "my error is:" + e.msg
}

func ReError() (error error) {
	return MyError{
		"sss",
	}
}
