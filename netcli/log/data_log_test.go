package log

import (
	"fmt"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	open, err := os.OpenFile("netcli.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	Log(open, "hello")
	Log(open, "hello2")

}
