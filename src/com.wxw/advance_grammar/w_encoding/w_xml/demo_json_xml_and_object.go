/*
@Time : 2022/1/5 10:45
@Author : wxw
@File : demo_xml_to_json
@Link: https://www.jianshu.com/p/28f57095fb14
*/
package main

import (
	"encoding/xml"
	"fmt"
)

type Students struct {
	XMLName xml.Name  `xml:"students"`
	Version string    `xml:"version,attr"`
	Stus    []Student `xml:"student"`
}

type Student struct {
	StudentID int    `xml:"student-id"`
	Name      string `xml:"student-name"`
	Address   string `xml:"address"`
	School    string `xml:"school"`
	Age       int    `xml:"age"`
	Gender    int    `xml:"gender"`
}

/**
 * 核心功能：xml 转json
 * 使用场景: 美团、微信
 */
func main() {
	fmt.Println("==================== 开始 ===================")
	// 编码
	CreateStudentXml()

	// 分隔符
	fmt.Println("==================== 分隔符 ==================")

	// 解码
	CreatStudentFromXml()
	fmt.Println("==================== 结束 ====================")
}

/**
 * 对象 转换为xml
 */
func CreateStudentXml() {
	stus := &Students{}
	stus.Version = "1"
	stus.Stus = append(stus.Stus, Student{
		StudentID: 3782378,
		Name:      "Java半颗糖",
		Address:   "北京市昌平区",
		School:    "TaiAn School",
		Age:       34,
		Gender:    0,
	})
	stus.Stus = append(stus.Stus, Student{
		StudentID: 3782379,
		Name:      "Java半颗糖",
		Address:   "北京市昌平区",
		School:    "ShiPai School",
		Age:       30,
		Gender:    1,
	})
	out, err := xml.Marshal(stus)

	if err != nil {
		fmt.Printf("error:%v\n", err)
	} else {
		//os.Stdout.Write([]byte(xml.Header))
		//os.Stdout.Write(out)
		fmt.Println(xml.Header + string(out))
	}
}

/**
 * xml 转换为 object
 */
func CreatStudentFromXml() {
	stus := Students{}
	datas := []byte(string(`
<?xml version="1.0" encoding="UTF-8"?>
    <students version="1">
         <student>
            <student-id>3782378</student-id>
            <student-name>江华</student-name>
            <address>北京昌平区</address>
            <school>TaiAn School</school>
            <age>34</age>
            <gender>0</gender>
        </student>
        <student>
            <student-id>3782379</student-id>
            <student-name>小歪子Go</student-name>
            <address>北京昌平区</address>
            <school>ShiPai School</school>
            <age>30</age>
            <gender>1</gender>
        </student>
    </students>
`))
	if err := xml.Unmarshal(datas, &stus); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stus)
}
