/*
@Time : 2022/4/20 16:39
@Author : weixiaowei
@File : init
*/
package etcds

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitClient() (*clientv3.Client, error) {
	// 初始化客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("err = ", err)
	}
	defer cli.Close()
	return cli, err
}
