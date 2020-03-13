package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
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

func demo1(db *gorm.DB) {
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{1, "七米", "男", "篮球"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu = new(UserInfo)
	db.Find(uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(u).Update("hobby", "a色球")
	fmt.Printf("%#v\n", u)
	// 删除
	db.Delete(u)
	db.Delete(&u2)
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	No1          *int    // 默认为null
	No2          int     // 默认为0
	Role         string  `gorm:"size:255"`       // 设置字段大小为255
	MemberNumber *string `gorm:"not null"`       // 设置会员号（member number）不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`              // 忽略本字段
}
type XUser struct {
	ID           int64
	Name         *string `gorm:"size:1023;default:'小王子';comment:'姓名'"`
	Age          *int64 `gorm:"not null;default:10"`
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func demo2(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	db.Create(&User{Name: "L1212", Role: "Admin", Email: "c"})
}

func demo3(db *gorm.DB) {
	db.AutoMigrate(XUser{})
	//user := XUser{Age: 18, Name: new(string)}

	//fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	//db.Create(&user)   // 创建user
	//fmt.Printf("%#v\n", user)
	//fmt.Println(db.NewRecord(user)) // 创建`user`后返回`false`
}
func main() {
	db := initDb()
	defer db.Close()
	//demo1(db)
	//demo2(db)
	demo3(db)
}
