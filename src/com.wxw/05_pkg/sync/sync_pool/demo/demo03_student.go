/*
@Time : 2022/3/29 10:46
@Author : weixiaowei
@File : demo03_student_benchmark
@相关文章：https://geektutu.com/post/hpg-sync-pool.html
*/
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// 定义一个池对象
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func main() {

	// 1.初始化student
	jsonObj := Student{
		Name: "Java半颗糖",
		Age:  int32(18),
	}
	bytesObj, _ := json.Marshal(jsonObj)
	// fmt.Println("student = ", string(bytesObj))

	// Get() 用于从对象池中获取对象，因为返回值是 interface{}，因此需要类型转换。
	// Put() 则是在对象使用完毕后，返回对象池。
	student := studentPool.Get().(*Student)
	if err := json.Unmarshal(bytesObj, student); err != nil {
		fmt.Println("student = ", student)
	}
	studentPool.Put(student)
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte `json:"remark"`
}
