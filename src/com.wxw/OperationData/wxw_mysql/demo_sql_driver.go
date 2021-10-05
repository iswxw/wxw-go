/*
@Time: 2021/10/5 20:48
@Author: wxw
@File: demo
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 基于 sql-driver 操作数据库
func main() {
	db, err := sql.Open("mysql", "root:123456@/wxw_test?charset=utf8")
	checkErr(err)

	//2. 插入数据
	stmt, err := db.Prepare("INSERT user_info SET username=?,company=?,createtime=?")
	checkErr(err)
	res, err := stmt.Exec("w02", "WeChat", "2012-12-09")
	checkErr(err)
	// 插入的Id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("插入的Id = %d\n", id)

	//3. 更新数据
	stmt, err = db.Prepare("update user_info set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("w02", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//1. 查询数据
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var company string
		var createtime string
		err = rows.Scan(&uid, &username, &company, &createtime)
		checkErr(err)
		fmt.Println(uid, username, company, createtime)
	}

	//删除数据
	stmt, err = db.Prepare("delete from user_info where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Printf("影响行数: %d\n", affect)

	// 关闭连接
	db.Close()

}

// 存在 error 进行 painc
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
