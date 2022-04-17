package models

import (
	"fmt"
	"webapp/cmdb/utils"
)

// User用户对象
type User struct {
	ID         int
	StaffID    string
	Name       string
	NickName   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	status     int
}

const (
	sqlQueryByName = "select id,name,password from user where name=?"
)

//
func GetUserByName(str string) *User {
	fmt.Println(str)
	user := &User{}
	if err := DB.QueryRow(sqlQueryByName, str).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		fmt.Println(user)
		return user
	}
	return nil
}

func (u *User) ValidPassword(password string) bool {
	fmt.Println(password, u.Password)
	return u.Password == utils.Md5Text(password)
}
