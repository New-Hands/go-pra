package log

import (
	"os"
)

// Log 数据记录
func Log(file *os.File, content string) {
	_, _ = file.WriteString(content + "\n")
}
