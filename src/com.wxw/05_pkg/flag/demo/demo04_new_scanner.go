/*
@Time : 2022/4/28 17:06
@Author : weixiaowei
@File : demo04_new_scanner
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 从标准输入 按行扫描
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}

	line := scanner.Text()
	fmt.Printf("bufio.NewScanner:%q\r\n", scanner.Text())
	data := strings.Split(line, " ")
	fmt.Println(data)
}
