/*
@Time: 2022/4/27 23:32
@Author: wxw
@File: demo01_os.args https://www.cnblogs.com/aaronthon/p/10883711.html
*/
package main

import (
	"fmt"
	"os"
)

//os.Args demo
func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

// 执行脚本
// go run demo01_args.go a b c d
// 或者 go build -o "demo01_args.go" 之后 ./demo01_args a b c d

//D:\Project\wxw-go\src\com.wxw\05_pkg\flag\demo>go run demo01_args.go a b c d
//args[0]=C:\Users\wxw\AppData\Local\Temp\go-build1027210589\b001\exe\demo01_args.exe
//args[1]=a
//args[2]=b
//args[3]=c
//args[4]=d
