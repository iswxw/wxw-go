// @Time : 2023/6/1 10:54
// @Author : xiaoweiwei
// @File : demo_reflect_test

package _4_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  rt := reflect.TypeOf(i)            //得到类型的元数据,通过t能获取类型定义里面的所有元素
  v := reflect.ValueOf(i)            //得到实际的值，通过v获取存储在里面的值，还可以去改变值
  tag := rt.Elem().Field(0).Tag      //获取定义在struct里面的标签
  name := v.Elem().Field(0).String() //获取存储在第一个字段里面的值
*/

func TestHelloWorld(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
