/*
@Time : 2022/3/28 10:47
@Author : weixiaowei
@File : demo_init
*/
package main

import (
	"context"
	"fmt"
	"framework/w_redis/init_redis"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

func main() {
	// 测试一：https://www.liwenzhou.com/posts/Go/go_redis/
	V8Example()
}

func V8Example() {
	ctx := context.Background()
	if err := init_redis.InitClient(); err != nil {
		return
	}
	err := init_redis.RDB.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := init_redis.RDB.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := init_redis.RDB.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
