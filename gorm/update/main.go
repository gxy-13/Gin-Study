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
		fmt.Println("链接数据库失败")
	} else {
		fmt.Println("数据库链接成功")
	}

	user := User{
		1, "Hello", "Gorm",
	}
	// update users set name = "Hello"  password = "Gorm" where id = 1
	db.Save(&user)

	//user := User{
	//	Username: "halo",
	//	Password: "halo",
	//}
	//db.Save(&user) // 如果不填写id 就会变成插入数据

	user2 := User{}
	db.First(&user2)
	fmt.Printf("result2 : %#v\n", user2)

	// 更新单个列 ，条件更新
	// update users set username = "halo2" where id = 9
	db.Model(&User{}).Where("id = ?", 9).Update("username", "halo2")

	// 更新多个列
	user3 := User{
		10, "testGorm", "ghj",
	}
	// 根据struct，只会更新非零值的字段
	// update users set username = "testGorm2" password = "YUI" where id = 10
	db.Model(&user3).Updates(User{Username: "testGorm2"})

	user4 := User{
		4, "ZZZ", "123",
	}
	// 根据map
	db.Model(&user4).Updates(map[string]interface{}{"username": "helloZzz", "password": "123123"})

	// 更新选择的字段
	user5 := User{
		5, "AAA", "12345",
	}
	db.Model(&user5).Select("password").Update("password", "123455")

	// 更新忽略字段
	user6 := User{
		6, "ccc", "123456",
	}
	db.Model(&user6).Omit("password").Update("username", "helloCcc")
}
