/*
@Time : 2022/4/20 16:17
@Author : weixiaowei
@File : demo_init_client
*/
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"log"
	"src/com.wxw/project_actual/src/03_thirdparty/w03_etcd/etcds"
)

func main() {
	// 初始化客户端
	cli, _ := etcds.InitClient()

	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 2000)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
		fmt.Println("err = ", err)
	}
	fmt.Println("resp = ", resp)

	// etcd 错误返回
	resp, err = cli.Put(ctx, "", "")
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	fmt.Println("resp = ", resp)
}
