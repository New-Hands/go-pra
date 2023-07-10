package pra_std

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	i, err := strconv.Atoi("66")

	if nil == err {
		fmt.Println(i)
	}
}
