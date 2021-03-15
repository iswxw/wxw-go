package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @Description
 * @Author wxw
 * @link
 * @Date 2021/3/15 22:02
 **/
func main() {
	// DSN:Data Source Name
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
