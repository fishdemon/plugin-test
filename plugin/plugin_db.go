package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

// gorm 默认的表明是结构体的复数 users
type User struct {
	ID string
	Name string
	Age int8
	Sex string
	Username string
	Password string
}

func init() {
	fmt.Println("init db plugin")
	db, err := gorm.Open("mysql", "root:123456@/go-study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed to init db")
	}
	//defer db.Close()

	// 禁用表名的复数形式
	db.SingularTable(true)
	Db = db
	fmt.Println("finish db plugin")
}

// 测试 curd
func TestCRUD()  {
	user := User{"12346", "john", 23, "M", "john","123456"}
	res := Db.Create(&user)
	if res.Error == nil {
		json1, _ := json.Marshal(user)
		fmt.Println(string(json1))
	}

	var user1 User
	res = Db.First(&user1, "12346")
	if res.Error == nil {
		json1, _ := json.Marshal(user1)
		fmt.Println(string(json1))
	}

	Db.Model(&user1).Update("age", 49)

	res = Db.First(&user1)
	if res.Error == nil {
		json1, _ := json.Marshal(user1)
		fmt.Println(string(json1))
	}

	Db.Delete(&user1)
}

func TestTx()  {
	// 开启事务
	tx := Db.Begin()

	//在事务中执行数据库操作，使用的是tx变量，不是db。
	user := User{"12346", "john", 23, "M", "john","123456"}
	err := tx.Create(&user).Error

	//保存订单失败，则回滚事务
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

