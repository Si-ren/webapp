package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //需要导入数据库所需 driver
	"time"
)

// beego的orm会找第一个int值设置为主键且自动增长
// 如果修改了列名,只会在数据库中加列,不会删
// 字段名小写不会映射,和json一样
type Student1 struct {
	ID         int    `orm:"column(id);pk;auto;"`
	Name       string `orm:"column(name1);size(16);description(姓名)" `
	Gender     bool   `orm:"default(true)"`
	Height     float32
	Rank       int        `orm:"index"`           //加上索引
	Birthday   *time.Time `orm:"type(date);null"` //默认时间类型为datetime,允许为null
	CreateTime *time.Time `orm:"auto_now_add"`    //自动添加当前时间
	UpdateTime *time.Time `orm:"auto_now"`        //默认为当前时间
}

//修改表名为stu,默认为student
func (s *Student1) TableName() string {
	return "stu"
}

//为表创建索引,可以为联合索引
func (s *Student1) TableIndex() [][]string {
	return [][]string{
		{"ID"},
		{"ID", "Name"},
		{"ID", "Name", "Rank"},
	}
}

func main() {
	//0.导入包
	//1.注册驱动
	//2.注册数据库
	//3.定义数据模型 model
	//4.注册数据模型
	//5.操作:
	//		同步表结构
	//		数据:增,删,改,查

	DSN := "root:root@tcp(localhost:3306)/cmdb?parseTime=true"
	//注册驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL) //orm已经注册了，可省略
	fmt.Println("orm.RegisterDriver", err)
	//注册数据库,必须有一个叫default的别名
	err = orm.RegisterDataBase("default", "mysql", DSN, 30)
	fmt.Println("orm.RegisterDataBase", err)

	//定义数据模型
	orm.RegisterModel(new(Student))

	orm.RunSyncdb("default", true, true)
	//开启命令行模式
	//syncdb     - auto create tables
	//sqlall     - print sql of create tables
	//help       - print this help
	// go run beego_orm.go  orm syncdb -force -v 会把表drop掉然后重新创建,所以这个只能在开发模式中使用

	orm.RunCommand()
}
