/*
 * @Time : 2021/3/3 20:33
 * @Author : wxw
 * @File : demo_json
 * @Software: GoLand
 * @Link: https://www.liwenzhou.com/posts/Go/13_reflect/
 * @Vlog: https://www.bilibili.com/video/BV1FV411r7m8?p=87
 */
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"小伟","age":9000}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

	// TypeOf
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64

}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	k := reflect.ValueOf(x)
	fmt.Printf("type:%v ,valueof:%v \n", v, k)
}
