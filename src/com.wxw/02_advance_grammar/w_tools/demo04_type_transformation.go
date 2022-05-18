/*
@Time : 2021/11/23 00:25
@Author : wxw
@File : demo_type_transformation 类型转换
*/
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/**
核心功能
 1. Json和struct互换
 2. json和map互转
 3. struct和map互转
*/

func main() {

	// 1. Json2Struct func Unmarshal(data []byte, v interface{}) error
	// 2. Struct2Json func Marshal(v interface{}) ([]byte, error)

	// 3. Json2Map()  func Unmarshal(data []byte, v interface{}) error
	// 4. Map2Json()  func Marshal(v interface{}) ([]byte, error)

	// 5. Struct2Map  func Struct2Map(obj interface{}) (mapStruct map[string]interface{},err error)
	// 6. Struct2MapByReflect func Struct2MapByReflect(obj interface{}) map[string]interface{}
	// 7. Map2Struct   方式1：通过第三方包 github.com/mitchellh/mapstructure 方式2：通过 map 转 json，再通过 json 转 struct

}

func Json2Map() {
	jsonStr := `{"name": "jqw","age": 18}`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("Json2Map err: ", err)
	}
	fmt.Println(mapResult)
}

func Map2Json() {
	var mapInstances []map[string]interface{}
	instance1 := map[string]interface{}{"name": "John", "age": 10}
	instance2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance1, instance2)
	jsonStr, err := json.Marshal(mapInstances)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

// Struct2Map 方式一
func Struct2Map(obj interface{}) (mapParams map[string]interface{}, err error) {
	var result map[string]interface{}
	if b, err := json.Marshal(obj); err == nil {
		if err := json.Unmarshal(b, &result); err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, err
}

// Struct2MapByReflect 通过反射实现 方式二
func Struct2MapByReflect(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Map2Struct() {

}
