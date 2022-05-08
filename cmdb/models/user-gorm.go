package models

//
//import (
//	"cmdb/utils"
//	"database/sql"
//	"fmt"
//	"golang.org/x/crypto/bcrypt"
//	"gorm.io/gorm"
//)
//
//// User 用户对象
//type User struct {
//	gorm.Model `gorm:"gorm.Model"`
//	//ID         int    `gorm:"index:idx_id;not null;autoIncrement"`
//	//string 类型 从gorm创建到数据库是varchar类型,可以写为 type:varcahr(32)
//	StaffID    string `gorm:"size:32"`
//	Name       string `gorm:"size:16"`
//	NickName   string `gorm:"size:16"`
//	Password   string `gorm:"size:128"`
//	Gender     int    `gorm:"type:tinyint"`
//	Tel        string `gorm:"size:16"`
//	Addr       string `gorm:"size:64"`
//	Email      string `gorm:"size:64"`
//	Department string `gorm:"size:16"`
//	//int无法设置size
//	Status int `gorm:"status"`
//}
//
//const (
//	sqlQueryByName = "select id,name,password from user.go where name=?"
//	sqlQuery       = "select gender, name,department from users"
//)
//
////
//func GetUserByName(name string) (*User, error) {
//	fmt.Println(name)
//	user.go := &User{}
//	//user1 := &User{}
//	//DB.Raw("SELECT * FROM `users` WHERE name='siri'").Scan(user1)
//	//fmt.Println(user1)
//	err := DB.Table("users").Where("name = ?", name).First(user.go).Error
//	if err == nil {
//		fmt.Println("Get user.go: ", user.go, err)
//		return user.go, err
//	}
//	fmt.Println("Can't get user.go: ", user.go, err)
//	return nil, err
//}
//
//func (u *User) ValidPassword(password string) bool {
//	fmt.Println(password, u.Password)
//	//fmt.Println(u.Password == utils.Md5Text(password))
//	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
//		return false
//	}
//	return true
//	//return u.Password == password
//}
//
//func CreateUser(user.go *User) (bool, error) {
//	user.go.Password = utils.Md5Text(user.go.Password)
//	err := DB.Create(user.go).Error
//	if err != nil {
//		return false, err
//	}
//	return true, err
//}
//
//func QueryUser(query string) ([]*User, error) {
//	users := make([]*User, 0)
//	var rows *sql.Rows
//	if query == "" {
//		//https://gorm.io/docs/sql_builder.html#Row-amp-Rows
//		rows, err = DB.Raw(sqlQuery).Rows()
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//	} else {
//		query = utils.Like(query)
//		fmt.Println(query)
//		SQL := sqlQuery + " where gender like ? ESCAPE '/' or  name like  ? ESCAPE '/' or department like ? ESCAPE '/'"
//		rows, err = DB.Raw(SQL, query, query, query).Rows()
//	}
//	defer rows.Close()
//	for rows.Next() {
//		user.go := &User{}
//		if err := rows.Scan(&user.go.Gender, &user.go.Name, &user.go.Department); err == nil {
//			users = append(users, user.go)
//		}
//	}
//	fmt.Println("User QueryUser :", users)
//	return users, err
//}
//
//// GenderText  性别显示
//func (u *User) GenderText() string {
//	if u.Gender == 0 {
//		return "女"
//	}
//	return "男"
//}
//
////StatusText 状态显示
//func (u *User) StatusText() string {
//	switch u.Status {
//	case 0:
//		return "正常"
//	case 1:
//		return "锁定"
//	case 2:
//		return "离职"
//	}
//	return "Error Status"
//}
