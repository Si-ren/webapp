package model

import (
	"fmt"
	"webapp/study/utils"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() {
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	inStmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常: ", err)
	}
	_, err = inStmt.Exec("admin", "123456", "siri@test.com")
	if err != nil {
		fmt.Println("inStmt.Exec err: ", err)
	}

	//_, err = utils.DB.Exec(sqlStr, "root", "root", "root@test.com")
	//if err != nil {
	//	fmt.Println("DB.Exec err: ", err)
	//}
}

//GetUserById 根据用户的id从数据库中查询一条记录
func (user *User) GetUserById() (*User, error) {
	//写sq1语句
	sqlStr := "select id,username , password, email from users where id = ?"
	//执行
	row := utils.DB.QueryRow(sqlStr, user.ID)
	//声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

//GetUsers获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//写sql语句
	sqlStr := "select id, username,password, email from users"
	//执行
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建切片
	var users []*User
	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, err
}
