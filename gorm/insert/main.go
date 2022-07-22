package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:13306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("链接数据库失败！")
	} else {
		fmt.Println("链接数据库成功")
	}

	user1 := User{
		1,
		"hh",
		"123",
	}
	// 通过数据的指针来创建
	result := db.Create(&user1)
	if result.RowsAffected == 1 {
		fmt.Println("数据插入成功")
	}
	user2 := User{
		2, "gxy", "12345",
	}

	// insert into `users` (`id`,`username`,`password`) values ("2","gxy","12345")
	db.Select("id", "username", "password").Create(&user2)

	user3 := User{
		3, "xxx", "123456",
	}

	//insert into `users` (username, password) values("xxx", "12312")
	// 可以用于主键自增的情况
	db.Omit("id").Create(&user3)

	// 批量插入,利用slice
	var users = []User{{Username: "ZZZ"}, {Username: "AAA"}, {Username: "ccc"}}
	db.Create(users)

	// 根据map 批量插入
	db.Model(&User{}).Create([]map[string]interface{}{
		{"username": "ggg", "password": "wqe"},
		{"username": "ggg", "password": "wqe"},
	})

}
