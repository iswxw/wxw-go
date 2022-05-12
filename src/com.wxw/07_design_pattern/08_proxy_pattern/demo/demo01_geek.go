/*
@Time : 2022/5/12 11:01
@Author : weixiaowei
@File : demo01_geek
*/
package main

import "fmt"

// StationProxy 代理了 Station，代理类中持有被代理类对象，并且和被代理类对象实现了同一接口。

// 售票员
type Seller interface {
	sell() // 售票
}

// 火车站
type Station struct {
	stock int // 余票库存
}

// 火车站售票
func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售罄")
	}
}

// 火车票代理站
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
	} else {
		fmt.Println("票已售罄")
	}
}
