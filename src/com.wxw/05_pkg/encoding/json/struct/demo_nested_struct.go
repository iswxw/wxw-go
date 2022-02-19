/*
@Time : 2022/2/12 20:57
@Author : weixiaowei
@File : demo_nested_struct
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 班级和学生
type (
	Class struct {
		ClassId   int      `json:"class_id"`
		ClassName string   `json:"class_name"`
		Student   *Student `json:"student"`
	}
	Student struct {
		StudentId   int    `json:"student_id"`
		StudentName string `json:"student_name"`
	}

	// 班级二
	Class1 struct {
		ClassId   int         `json:"class_id"`
		ClassName string      `json:"class_name"`
		Student   interface{} `json:"student"`
	}
)

// 嵌套结构体
func main() {

	fmt.Println("============结构体中包含结构体==============")
	TestNestedStruct()

	fmt.Println("============结构体中包含接口==============")

	// 结构体中包含接口
	jsonStr := `{"class_id":1,"class_name":"子材班","student":{"student_id":11,"student_name":"魏同学"}}`
	class1 := &Class1{}
	if err := json.Unmarshal([]byte(jsonStr), class1); err != nil {
		log.Println("err:", err)
	}
	log.Printf("Unmarshal:%#v", class1)

	marshal, err := json.Marshal(class1.Student)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("Marshal:", string(marshal))

}

// 嵌套结构通序列化和反序列化
func TestNestedStruct() {
	// 序列化
	class := Class{
		ClassId:   1,
		ClassName: "子材班",
		Student: &Student{
			StudentId:   11,
			StudentName: "魏同学",
		},
	}
	byteClass, err := json.Marshal(class)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("Marshal:", string(byteClass))

	// 反序列化
	newClass := Class{}
	if err = json.Unmarshal(byteClass, &newClass); err != nil {
		log.Println("err:", err)
	}
	log.Printf("Unmarshal:%#v", newClass)

	marshal, err := json.Marshal(newClass)
	log.Println("Unmarshal:", string(marshal))
}
