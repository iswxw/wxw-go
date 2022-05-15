/*
@Time: 2021/10/5 21:23
@Author: wxw
@File: demo_beego
*/
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// UserInfo model struct
type UserInfo struct {
	Uid        int64 `orm:"colunm(uid);pk"` // 设置主键
	Username   string
	Company    string
	Createtime string `default:"1994-01-01"` // 设置默认值
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/wxw_test?charset=utf8", 30)
	// 注册定义的 model 可以同时注册多个 model
	orm.RegisterModel(new(UserInfo))
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 如果不存在就创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	fmt.Println("beego开始操作...")
	// 基本的赋值
	user := UserInfo{Username: "slene", Createtime: "2021-08-12"}

	// 打印记录
	fmt.Println(user.Uid, user.Username, user.Company, user.Createtime)

	// 开启一个连接
	o := orm.NewOrm()

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	//user.Username = "w4"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := UserInfo{Uid: user.Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	//
	//	// 删除表
	//	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
