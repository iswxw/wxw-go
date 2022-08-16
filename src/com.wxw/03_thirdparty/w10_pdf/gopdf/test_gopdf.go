// @Time : 2022/8/16 14:43
// @Author : xiaoweiwei
// @File : gopdf_test

package main

import (
	"log"
	"os"
)

func main() {
	CreateFile()
}

func CreateFile() {
	create, err := os.Create("./test.txt")
	defer create.Close()
	if err != nil {
		log.Println("err = ", err)
	}
	create.WriteString("hello wxw-go!")
}
