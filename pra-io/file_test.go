package pra_io

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, err := os.ReadFile("C:\\Users\\lst\\Desktop\\infoboard\\play-vs2.lst")
	if err != nil {
		return
	}

	fmt.Println(file)
	dump := hex.EncodeToString(file)
	fmt.Println(dump)
}

// 创建文件
func TestCreate(t *testing.T) {

}
