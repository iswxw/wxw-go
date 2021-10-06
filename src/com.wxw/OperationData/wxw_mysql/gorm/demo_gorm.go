/*
@Time: 2021/10/6 13:14
@Author: wxw
@File: demo_gorm
*/
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type UserInfo struct {
	Uid        int64 `gorm:"primaryKey"`
	UserName   string
	Company    string
	CreateTime time.Time
}

var db *gorm.DB

func main() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/wxw_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		panic(err)
	}

	// 定义指针
	var userinfo UserInfo

	// 获取第一条记录
	// 有效，因为目标 struct 是指针
	// SELECT * FROM `user_info` ORDER BY `user_info`.`uid` LIMIT 1
	db.First(&userinfo)
	fmt.Println(userinfo)
}
