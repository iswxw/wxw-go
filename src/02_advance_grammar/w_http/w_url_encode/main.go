/*
@Time : 2022/1/20 23:41
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"net/url"
)

// 1. url.Encode() 函数会将每个参数进行url编码，然后参数之间用&符号连接，对应的key和value用'='连接
// 2. url.QueryUnescape() 将转码的数据解码 恢复正常
func main() {
	data := url.Values{}
	data.Set("uid", "1")
	data.Set("phone", "2")
	data.Set("name", "liyunlong")
	data.Set("data", "http://alipared:10003/paydcaldlback/alidpay/432412/Nofaffa")

	fmt.Println(data.Encode())
	//tmp := url.Values{"data":[]{"data is nklsjfklajf"}}
	fmt.Println(url.QueryUnescape(data.Encode()))
}

// 输出
// data=http%3A%2F%2Falipared%3A10003%2Fpaydcaldlback%2Falidpay%2F432412%2FNofaffa&name=liyunlong&phone=2&uid=1
// data=http://alipared:10003/paydcaldlback/alidpay/432412/Nofaffa&name=liyunlong&phone=2&uid=1 <nil>
