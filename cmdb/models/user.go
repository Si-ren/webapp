package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户对象
type User struct {
	ID         int    `orm:"column(id)"`
	StaffID    string `orm:"size(32)"`
	Name       string `orm:"size(32)"`
	NickName   string `orm:"size(32)"`
	Password   string `orm:"size(1024)"`
	Gender     int    `orm:"type(tinyint)"`
	Tel        string `orm:"size(16)"`
	Addr       string `orm:"size(128)"`
	Email      string `orm:"size(128)"`
	Department string `orm:"size(32)"`
	//int无法设置size
	Status   int       `orm:""`
	CreateAt time.Time `orm:"auto_now_add"`
	UpdateAt time.Time `orm:"auto_now"`
	DeleteAt time.Time `orm:"null "`
}

const (
	sqlQueryByName = "select id,name,password from user.go where name=?"
	sqlQuery       = "select gender, name,department from users"
)

//
//func GetUserByName(name string) (*User, error) {
//	fmt.Println(name)
//	user := &User{}
//	user.Name = name
//
//	if err := mysql.Read(user, "name"); err != nil {
//		return nil, err
//	} else {
//		return user, err
//	}
//}

func (u *User) ValidPassword(password string) bool {
	fmt.Println(password, u.Password)
	//fmt.Println(u.Password == utils.Md5Text(password))

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		fmt.Println("ValidPassword bcrypt.CompareHashAndPassword :", err)
		return false
	}
	return true
	//return u.Password == password
}

//func CreateUser(user *User) (bool, error) {
//	//user.go.Password = utils.Md5Text(user.go.Password)
//	fmt.Println(user.Password)
//	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
//	if err != nil {
//		return false, err
//	}
//
//	user.Password = string(password)
//	fmt.Println(user.Password)
//
//	_, err = mysql.Insert(user)
//	if err != nil {
//		return false, err
//	}
//	return true, err
//}

// GenderText  性别显示
func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}

//StatusText 状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "Error Status"
}

//func DeleteUserByID(ID int) {
//	mysql.Delete(&User{ID: ID})
//}
//
//func ModifyUserPassword(user *User, password string) {
//	user.Password = utils.GeneratePassword(password)
//	mysql.Update(user, "Password")
//}
