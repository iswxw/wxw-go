/*
@Time : 2022/4/20 16:31
@Author : weixiaowei
@File : demo02_crud
*/
package main

import (
	"context"
	"fmt"
	"framework/w_etcd/etcds"
)

func main() {
	cli02, _ := etcds.InitClient()
	ctx := context.Background()
	putResponse, _ := cli02.Put(ctx, "key", "value")
	fmt.Println("putResponse = ", putResponse)

	getResponse, _ := cli02.Get(ctx, "key")
	fmt.Println("getResponse = ", getResponse)

	return
}
