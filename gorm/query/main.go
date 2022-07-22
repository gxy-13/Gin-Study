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

	user := User{}
	// 获取第一条记录(主键升序) select * from users order by id limit 1
	result := db.First(&user)
	if result.RowsAffected == 1 {
		fmt.Printf("result1 : %#v\n", user)
	} else {
		fmt.Println(result.Error)
	}

	// 获取第一条记录，没有指定排序字段 select * from users limit 1
	result2 := db.Take(&user)
	if result2.RowsAffected == 1 {
		fmt.Printf("result2 : %#v\n", user)
	} else {
		fmt.Println(result2.Error)
	}

	// 获取最后一条记录(主键降序)  select * from users order by id desc limit 1
	result3 := db.Last(&user)
	if result3.RowsAffected == 1 {
		fmt.Printf("result3 : %#v\n", user)
	} else {
		fmt.Println(result3.Error)
	}

	// first 和 last 只有目标struct是指针或者通过db.model() 指定model时，该方法才有效
	// 如果model没有定义主键，那么对model的第一个字段进行排序
	result4 := map[string]interface{}{}
	db.Model(&User{}).First(&result4) // select * from users order by id limit 1
	fmt.Printf("result4 : %#v\n", result4)

	// 用主键搜索 select * from users where id = 10
	db.First(&user, 10)
	fmt.Printf("result5: %#v\n", user)

	users := make([]User, 0, 3)
	// select * from users where id in (1,2,3)
	db.Find(&users, []int{1, 2, 3})
	fmt.Printf("result6: %#v\n", user)

	// 如果主键时字符串，例如uuid select * from users where id = "1b74413f-f3b8-409f-ac47-e8c062e3472a"
	db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	fmt.Printf("result7 :%#v\n", user)

	// 条件查询 string
	db.Where("username = ?", "gxy").First(&user)
	fmt.Printf("result 8 :%#v\n", user)

	// 条件查询 map 或者 struct select * from users where username = "gxy" and password = "12345"
	db.Where(&User{Username: "gxy", Password: "12345"}).Find(&user)
	fmt.Printf("result9 :%#v\n", user)

	// 高级查询 假设 User中除了id ，username，password 之外还有几百个字段
	// 我们可以定义一个较小的结构体，来实现调用api时自动选择特定的字段
	type APIUser struct {
		ID       int
		Username string
	}
	var apiUser []APIUser
	// 查询时会自动选择id ， name 字段 select id, name from users limit 10
	db.Model(&User{}).Limit(10).Find(&apiUser)
	for _, u := range apiUser {
		fmt.Printf("result10 :%#v\n", u)
	}

}
