package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func demo(db *gorm.DB) {
	db.Exec("drop table if exists user")
	db.AutoMigrate(&User{})
	user := User{Name: sql.NullString{"", false}, Age: 18} // 此时数据库中该条记录name字段的值就是'小王子'
	//user := User{Name: sql.NullString{"", true}, Age: 18}  // 此时数据库中该条记录name字段的值就是''
	db.Create(&user)
	fmt.Printf("%#v\n", user.Name.String)

	user = User{Name: sql.NullString{"", true}, Age: 20}
	db.Create(&user)

	// 根据主键查询第一条记录
	db.First(&user)
	db.First(&User{})
	db.Last(&User{})

	// 下面的语句没有区别
	users := []User{}
	db.Find(&users)
	fmt.Printf("%#v\n", users)

	users = []User{}
	db.Where("name = ?", "小王子").First(&User{})
	db.Where("name = ?", "小王子").Find(&users)
	fmt.Printf("%#v\n", users)

}

type User struct {
	gorm.Model
	Name sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
	Age  int64
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/tb2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 自动迁移
	db.LogMode(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	// 全局禁用表名复数
	db.SingularTable(true)
	return db
}

func main() {
	db := initDb()
	defer db.Close()
	demo(db)
}
