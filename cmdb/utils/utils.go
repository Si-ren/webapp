package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Text(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
