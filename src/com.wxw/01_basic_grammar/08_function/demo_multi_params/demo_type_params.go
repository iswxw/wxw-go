/*
@Time: 2022/4/10 17:50
@Author: wxw
@File: demo_type_params
*/
package main

import "fmt"

// 函数式编程实现
// https://studygolang.com/articles/34446
var defaultStuffClient = stuffClient{
	retries: 3,
	timeout: 2,
}

type StuffClientOption func(*stuffClient)

func WithRetries(r int) StuffClientOption {
	return func(o *stuffClient) {
		o.retries = r
	}
}
func WithTimeout(t int) StuffClientOption {
	return func(o *stuffClient) {
		o.timeout = t
	}
}

type StuffClient interface {
	DoStuff() error
}

type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}
type Connection struct{}

// 新建一个客户端
func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	client := defaultStuffClient
	for _, o := range opts {
		o(&client)
	}
	client.conn = conn
	return client
}
func (c stuffClient) DoStuff() error {
	return nil
}

func main() {
	x := NewStuffClient(Connection{})
	fmt.Println(x) // prints &{{} 2 3}

	x = NewStuffClient(Connection{}, WithRetries(1))
	fmt.Println(x) // prints &{{} 2 1}

	// 选择参数传递，并且是指定的参数
	x = NewStuffClient(Connection{}, WithRetries(1), WithTimeout(1))
	fmt.Println(x) // prints &{{} 1 1}
}
