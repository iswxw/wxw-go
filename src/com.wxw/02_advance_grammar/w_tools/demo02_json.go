/*
@Time : 2022/5/18 15:49
@Author : weixiaowei
@File : demo_json
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Person02 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person02{
		Name: "Java 半颗糖",
		Age:  18,
	}
	pretty := jsonPretty(person)
	fmt.Println(pretty)
}

func jsonPretty(d interface{}) string {
	r, _ := json.MarshalIndent(d, "", "\t")
	return string(r)
}
