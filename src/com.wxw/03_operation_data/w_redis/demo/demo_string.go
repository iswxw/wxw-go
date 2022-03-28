/*
@Time : 2022/3/28 11:02
@Author : weixiaowei
@File : demo_string
*/
package main

import (
	"context"
	"fmt"
	"framework/w_redis/init_redis"
	"github.com/go-redis/redis/v8"
)

func main() {

	redisExample()
}

func redisExample() {
	ctx := context.Background()
	if err := init_redis.InitClient(); err != nil {
		return
	}

	err := init_redis.RDB.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := init_redis.RDB.Get(ctx, "score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := init_redis.RDB.Get(ctx, "name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
