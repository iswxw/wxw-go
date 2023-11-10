/*
 * @Time : 2021/3/3 22:03
 * @Author : wxw
 * @File : demo_parse_ini
 * @Software: GoLand
 * @Link:
 * @Vlog: https://www.bilibili.com/video/BV1FV411r7m8?p=91
 */
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// ini 配置文件解析器

// mysql 配置解析 结构体
type MySQLConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// redis 配置解析 结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// 结构体嵌套
type Config struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"mysql"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0.参数校验
	//0.1 入参必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Printf("type: %v, kind: %v \n ", t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer.") // 新创建一个错误
		return
	}
	//0.2 入参必须是结构体类型（因为配置文件中各种键值对需要赋值给结构体） t.Elem() 判断指针的类型
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer.") // 新创建一个错误
		return
	}
	//1. 读文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// string(b) // 将字符类型的文件内容转换为字符村
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	//2. 一行一行的取数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾空格
		line = strings.TrimSpace(line)
		//2.1 如果是注释则跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2 如果是[ 开头则是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			// 把这一行首尾的[] 去掉，去到中间的内容，把首尾空格去掉，取到内容
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			// 根据字符串section 去data里面根据反射找对应的结构体
			// v := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到 %s 对应嵌套的结构体：%s \n", sectionName, structName)
				}
			}
		} else {
			//2.3 如果不是[开头就是 = 分割的键值对

		}

	}
	return err
}

func main() {
	// fmt.Println(os.Getwd())
	var cfg Config
	err := loadIni("./src/com.wxw/01_basic_grammar/06_go_ini/conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err: %v \n", err)
		return
	}
	fmt.Println(cfg)
}
