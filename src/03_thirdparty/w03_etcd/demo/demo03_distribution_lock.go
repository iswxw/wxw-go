/*
@Time : 2022/4/20 17:11
@Author : weixiaowei
@File : demo03_distribution_lock
*/
package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

/**
来源：http://blueskykong.com/2021/05/06/etcd-opt/
总得来说，如上关于 etcd 分布式锁的实现过程分为四个步骤：
  1. 客户端初始化与建立连接；
  2. 创建租约，自动续租；
  3. 创建事务，获取锁；
  4. 执行业务逻辑，最后释放锁。
*/

func main() {
	// 客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1. 上锁并创建租约
	lease := clientv3.NewLease(client)
	leaseGrantResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		panic(any(err))
	}
	leaseId := leaseGrantResp.ID

	// 2 自动续约
	// 创建一个可取消的租约，主要是为了退出的时候能够释放
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 3. 释放租约
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)
	keepRespChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		panic(any(err))
	}
	// 续约应答
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else { // 每秒会续租一次, 所以就会受到一次应答
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 1.3 在租约时间内去抢锁（etcd 里面的锁就是一个 key）
	kv := clientv3.NewKV(client)

	// 创建事务
	txn := kv.Txn(context.TODO())

	//if 不存在 key，then 设置它，else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("lock"), "=", 0)).
		Then(clientv3.OpPut("lock", "g", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("lock"))

	// 提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		panic(any(err))
	}

	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 抢到锁后执行业务逻辑，没有抢到退出
	fmt.Println("开始处理任务")
	time.Sleep(10 * time.Second)

	fmt.Println("处理任务完成")
}
