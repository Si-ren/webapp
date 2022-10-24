package utils

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Md5Text(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//-Like sql查询字符串前后加%
func Like(q string) string {
	q = strings.TrimSpace(q)
	if q == "" {
		return ""
	}
	q = strings.Replace(q, "/", "//", -1)
	q = strings.Replace(q, "%", "/%", -1)
	q = strings.Replace(q, "_", "/_", -1)
	return fmt.Sprintf("%%%s%%", q)
}

func GeneratePassword(password string) string {
	tempPWD, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		fmt.Println(err)
	}
	return string(tempPWD)
}
