/*
@Time : 2022/3/29 20:03
@Author : weixiaowei
@File : demo_update_field
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	config := Config{}
	byteConfig01, err := json.Marshal(config)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("byteConfig = %#v \n", config)
	fmt.Println("byteConfig = ", string(byteConfig01))

	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "iswxw.gitee.io")

	c := readConfig()
	fmt.Printf("byteConfig = %#v \n", c)
}

func readConfig() *Config {
	// read from xxx.json，省略
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}
