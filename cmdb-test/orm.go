package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// Object Relational Mapping
type Product struct {
	gorm.Model
	Code  string `gorm:"index:idx_id;index:idx_oid,unique;size:32;not null;default:123"` //创建列的一些属性
	Price uint   `gorm:"column:pricesss"`                                                //修改表名
}

// 修改表名,不然默认为小写的结构体名称
func (p *Product) TableName() string {
	return "prod"
}

func main() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb-test",
		"root", "root", "127.0.0.1", "3306",
	)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		//定义log level: log levels: Silent, Error, Warn, Info
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
		os.Exit(-1)
	}

	// Migrate the schema
	//自动迁移只会添加列和索引,不会修改列类型和删除列
	db.AutoMigrate(&Product{})

	//判断表是否存在
	fmt.Println(db.Migrator().HasTable(&Product{}))
	// Create 行
	//插入日期: time.Date(1996,6,24,0,0,0,0,time.UTC)
	db.Create(&Product{Code: "D42", Price: 100})
	//如果没有数据,那么就插入
	if db.Find(&Product{Price: 300}).Error != nil {
		db.Create(&Product{Price: 300})
	}
	var prod Product
	if err := db.Where("pricesss = ?", 400).First(&prod).Error; err != nil {
		fmt.Println(err)
		fmt.Println(prod)
	}

	// Select
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//按条件在表中查找
	//https://gorm.io/docs/query.html
	//Not Or And and other conditions
	//db.Where("price=?", "200").Find(&Product{})
	//db.Where("price=?", []int{100, 200}).Find(&Product{})

	// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//逻辑删除
	//db.Delete(&product, 1)
	//真正删除
	//db.Unscoped().Delete(&product)
	//列
	//更改列名
	//db.Migrator().RenameColumn(&Product{}, "price", "prices")
	//删除列
	//db.Migrator().DropColumn(&Product{}, "price")

	//删除索引
	//db.Migrator().DropIndex(&Product{}, "idx_oid")

	//time.Sleep(9000 * time.Millisecond)
	//drop table
	//db.Migrator().DropTable(&Product{})
}
