/*
@Time : 2021/10/13 15:15
@Author : wxw
@File : demo_json
@link：
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type Student struct {
	ID     interface{}
	Name   string
	Sex    string
	Course []string
}

func main() {
	fmt.Println("==================== 开始 ==================")
	// 编码
	EncodeStruct()

	// 分隔符
	fmt.Println("==================== 分隔符 ==================")

	// 解码
	DecodeJsonString()
	fmt.Println("==================== 结束 ==================")
}

/**
 * 将一个对象编码成json字符串
 */
func EncodeStruct() {
	student := Student{
		ID:     14343466,
		Name:   "Java半颗糖",
		Sex:    "M",
		Course: []string{"English", "Math", "Chinese"},
	}
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

/**
 * 将json字符串解码成json对象
 */
func DecodeJsonString() {
	var jsonBody = []byte(`{"ID":14343466,"Name":"Java半颗糖","Sex":"M","Course":["English","Math","Chinese"]}`)
	var student1 = Student{}
	err := json.Unmarshal(jsonBody, &student1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(student1)
	fmt.Println("studentName:", student1.Name)
	fmt.Println("studentId:", student1.ID)
	fmt.Println("studentId_int:", cast.ToInt64(student1.ID))
	fmt.Println("studentSex:", student1.Sex)
	fmt.Println("studentCourse:", student1.Course)
}
