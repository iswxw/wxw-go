/*
@Time: 2021/10/7 0:43
@Author: wxw
@File: demo_tag
*/
package _struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// Teacher 学生
type Teacher struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问

}

// TestUsedTag 使用tag
func TestUsedTag(t *testing.T) {
	s1 := Teacher{
		ID:     1,
		Gender: "男",
		name:   "半颗糖",
	}
	// 序列化
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}

// TestCustomTag 自定义tag，详见：https://www.yuque.com/fcant/go/mhbard
func TestCustomTag(t *testing.T) {
	u := TagUser{
		Name:     "技术能量站",
		Age:      5,
		Password: "root",
	}
	rt := reflect.TypeOf(u)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get("wxw")
		fmt.Println("get tag is ", tag)
	}
}

type TagUser struct {
	Name     string `wxw:"Username"`
	Age      uint16 `wxw:"age"`
	Password string `wxw:"min=6,max=10"`
}
