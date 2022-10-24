package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //需要导入数据库所需 driver
	"time"
)

// beego的orm会找第一个int值设置为主键且自动增长
// 如果修改了列名,只会在数据库中加列,不会删
// 字段名小写不会映射,和json一样   如果使用驼峰式的命名规则,则会使用下划线分割开,例如: StudentRole : student_role
type Student struct {
	ID         int    `orm:"column(id);pk;auto;"`
	Name       string `orm:"column(name1);size(16);description(姓名)" `
	Gender     bool   `orm:"default(true)"`
	Height     float32
	Rank       int        `orm:"index"`           //加上索引
	Birthday   *time.Time `orm:"type(date);null"` //默认时间类型为datetime,允许为null
	CreateTime *time.Time `orm:"auto_now_add"`    //自动添加当前时间
	UpdateTime *time.Time `orm:"auto_now"`        //默认为当前时间
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

	//开启debug模式
	orm.Debug = true
	DSN := "root:root@tcp(localhost:3306)/cmdb-test?parseTime=true"
	//注册驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL) //orm已经注册了，可省略
	fmt.Println("orm.RegisterDriver", err)
	//注册数据库,必须有一个叫default的别名
	err = orm.RegisterDataBase("default", "mysql", DSN, 30)
	fmt.Println("orm.RegisterDataBase", err)

	orm.RegisterModel(new(Student))

	//建立数据库连接
	mysql := orm.NewOrm()
	//使用别的数据库,需要建立数据库连接后,使用Using指定数据库
	//mysql.Using("mysqlxxx")
	//mysql.Insert(xxx)

	/*
		//定义数据模型
		birthday, _ := time.Parse("2006-01-02", "1996-01-01")
		stu1 := &Student{
			ID:       0,
			Name:     "siri",
			Gender:   false,
			Height:   171,
			Rank:     1,
			Birthday: &birthday,
		}

		fmt.Printf("%#v\n", stu1)
		id, err := mysql.Insert(stu1)
		fmt.Println(id, err)
		//在插入后, beego/orm自动把标签设置为auto_now的列的时间加上了
		fmt.Printf("%#v\n", stu1) //&main.Student{ID:2, Name:"siri", Gender:false, Height:171, Rank:1, Birthday:time
		//.Date(1996, time.January, 1, 0, 0, 0, 0, time.UTC), CreateTime:time.Date(2022, t
		//ime.April, 26, 2, 51, 58, 986782400, time.Local), UpdateTime:time.Date(2022, tim
		//e.April, 26, 2, 51, 58, 986782400, time.Local)}

		users := make([]*Student, 0)
		for i := 0; i < 10; i++ {
			users = append(users, &Student{
				Name:       fmt.Sprintf("siri-%d", i),
				Gender:     false,
				Height:     float32(i),
				Rank:       i,
				Birthday:   nil,
				CreateTime: nil,
				UpdateTime: nil,
			})
		}
		//批量插入
		mysql.InsertMulti(3, users)
		//orm.RunSyncdb("default", false, true)
	*/

	/*
		//读取 select
		stu2 := &Student{ID: 1}
		err = mysql.Read(stu2)
		fmt.Println(" mysql.Read: ", stu2, err)

		//按照条件查找,如果满足条件的有多个,那么只会返回1条数据,所以需要保证数据是唯一的才准确
		stu3 := &Student{Name: "siri-1", Gender: false}
		err = mysql.Read(stu3, "Name", "Gender")
		fmt.Println(" mysql.Read: ", stu3, err)
	*/

	/*
		//update 更新
		stu4 := &Student{ID: 1}
		stu4.Name = "LSL"
		mysql.Update(stu4)
		//只更新Name字段
		mysql.Update(stu4,"Name")
	*/

	/*
		//delete 删除操作
		fmt.Println(mysql.Delete(&Student{ID: 2}))
	*/

	/*
		//不存在才创建
		stu4 := &Student{ID: 2}
		//如果ID为2的数据不存在,那么就创建
		mysql.ReadOrCreate(stu4, "ID")
	*/

	/*
		//查找
		//查找表,以下两种是一样的效果
		//mysql.QueryTable("stu")
		querySet := mysql.QueryTable(&Student{})
		fmt.Println(querySet.Count())
		var stus []*Student
		//查询所有
		querySet.All(&stus)
		fmt.Println(stus)
		//where Filter,格式: 列名__条件,对象
		//
		//		条件
		//	    = ,!= : exact , iexact
		//		< ,> : lt , gt
		//		<= , >= : lte , gte
		//		%content% : [i]contains
		//		start% : [i]startwith
		//		%end : [i]endwith
		//		in : in  ,对象要写切片
		//
		querySet.Filter("name__exact", "siri").Count()
		querySet.Filter("id__in", []int{3, 4, 5}).Filter("id__gt", 3).Count()
		//分页
		querySet.Limit(3).Offset(2).All(&Student{})
		//排序  降序列名前加个"-"
		querySet.OrderBy("Name").All(&Student{})
		querySet.OrderBy("-Name", "Gender").All(&Student{})
		//查询一条数据
		querySet.OrderBy("-Name", "Gender").One(&Student{})
		//查询一条数据,只获取id
		querySet.OrderBy("-Name", "Gender").One(&Student{},"id")

		//复杂条件查询
		//创建条件
		cond := orm.NewCondition()
		conHeight := orm.NewCondition()
		//制定条件
		conHeight = conHeight.And("height__lt", "200").And("height__gt", "150")
		cond = cond.OrCond(conHeight).And("name__in", []string{"siri", "lsl"})
		//使用条件  SetCond
		querySet.SetCond(cond).All(&Student{})

		//批量按条件更新
		querySet.Filter("id__gt", "10").Update(orm.Params{"name": "lsl", "height": orm.ColValue(orm.ColAdd, 10)})
	*/

	//执行sql方式一
	//db, err := orm.GetDB("default")
	//db.Exec()
	//db.QueryRow().Scan()

	//执行sql方式二
	//rawseter := mysql.Raw("INSERT INTO student (name1,gender,height,`rank`,create_time,update_time)  VALUES ('lsl123',0,158,'1','2022-04-25 19:00:45','2022-04-25 19:00:45')")
	//fmt.Println(rawseter.Exec())

	//通过sql获取值
	//var stus []*Student
	//rawseter := mysql.Raw("select id , name1 from student")
	//rawseter.QueryRows(&stus)
	//fmt.Println(stus)
	//fmt.Printf("%#v\n", stus[0])

	//通过sql获取sql函数值
	rawseter := mysql.Raw("select name1 ,count(*) from student  group by name1")
	//orm.Params 类型是key 和 value的结果
	var result []orm.Params
	rawseter.Values(&result)
	fmt.Println(result)
	//orm.ParamsList 类型是 切片
	var result1 []orm.ParamsList
	rawseter.ValuesList(&result1)
	fmt.Println(result1)

	orm.RunCommand()

}
