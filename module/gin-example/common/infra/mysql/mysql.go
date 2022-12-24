/*
@Time: 2022/12/24 22:39
@Author: wxw
@File: mysql
*/
package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"src/com.wxw/project_actual/module/gin-example/common/infra/conf"
)

// client gorm DB 的对象是线程安全的
var client *gorm.DB

// Setup 设置mysql启动项
func Setup() {
	var err error
	dsn := conf.Viper.GetString("mysql.dsn")
	client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		panic(fmt.Sprintf(" setup mysql failed: %s", err))
	}

	log.Printf("[mysql] setup success;\n")
}
