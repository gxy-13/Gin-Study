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
	dsn := "root:root@(127.0.0.1:13306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("链接数据库失败")
	} else {
		fmt.Println("数据库连接成功")
	}

	user1 := User{
		9, "hh", "zz",
	}
	// gorm 将根据主键删除记录 delete from users where id = 9
	db.Delete(&user1)

	// 批量删除 delete from users where username like %hello%
	db.Where("username like ?", "%hello%").Delete(User{})
}
