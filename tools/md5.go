package tools

import (
	"crypto/md5"
	"fmt"
)

func Md5Text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text))) //密文工具
}
