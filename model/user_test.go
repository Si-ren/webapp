package model

import (
	"fmt"
	"testing"
)

//被测试源文件必须和测试文件在同一目录下
//测试文件必须以_test.go结尾
//测试文件中的测试函数为TestXXX(t *testing.T),其中XXX的首字母必须大写

//TestMain 是在测试之前执行一些代码,做初始化的用处
func TestMain(m *testing.M) {
	fmt.Println("开始初始化")
	fmt.Println("开始测试")
	fmt.Println("开始连接")
	m.Run()
}

func TestUser_AddUser(t *testing.T) {
	fmt.Println("测试添加用户：")
	user := &User{}
	user.AddUser()
}

//函数名不以Test开头,函数默认不执行,可以设置为一个子测试函数
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的方法")
	//t.Run("测试添加用户：", testUser_AddUser2)
	//t.Run("测试查询用户：", testGetUserById)
	t.Run("测试查询用户：", testGetUsers)
}

func testUser_AddUser2(t *testing.T) {
	fmt.Println("测试添加用户：")
	user := &User{}
	user.AddUser()
}

func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录: ")
	user := &User{
		ID: 1,
	}
	//调用获取User的方法
	u, _ := user.GetUserById()
	fmt.Println("User的信息是:", u)
}

//测试获取所有User
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录: ")
	user := &User{}
	//调用获取所有User的方法
	us, _ := user.GetUsers()
	//遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户是%v: \n", k+1, v)
	}
}
