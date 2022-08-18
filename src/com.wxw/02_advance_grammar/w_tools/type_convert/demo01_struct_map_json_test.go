// @Time : 2022/8/18 15:53
// @Author : xiaoweiwei
// @File : demo01_struct_map_json_test

package type_convert

import (
	"log"
	"testing"
)

//  Struct2Map
//方式1：通过第三方包 github.com/mitchellh/mapstructure
//方式2：通过 map 转 json，再通过 json 转 struct
func TestStruct2Map(t *testing.T) {

	result, err := Struct2Map(Person{
		Name: "IsWxw",
		Age:  18,
	})

	//result := Struct2MapByReflect(Person{
	//	Name: "IsWxw",
	//	Age:  18,
	//})
	log.Println("err = ", err)
	log.Printf("result = %v \n", result)
}
