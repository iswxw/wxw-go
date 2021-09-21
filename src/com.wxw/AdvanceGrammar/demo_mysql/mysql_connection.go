package main

import (
	"database/sql"
)

/**
 * @Description
 * @Author wxw
 * @link
 * @Date 2021/3/15 22:02
 **/
func main() {
	// DSN:Data Source Name
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
