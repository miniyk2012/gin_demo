package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func demo(db *gorm.DB) {
	db.Exec("drop table if exists user")
	db.AutoMigrate(&User{})
	user := User{Name: new(string), Age: 18}
	db.Create(&user) // 此时数据库中该条记录name字段的值就是''
	fmt.Printf("%#v\n", user)

	user = User{Age: 18}
	db.Create(&user)
	fmt.Printf("%#v\n", *user.Name)
}

type User struct {
	ID   int64
	Name *string `gorm:"default:'小王子'"`
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
	//demo1(db)
	//demo2(db)
	demo(db)
}
