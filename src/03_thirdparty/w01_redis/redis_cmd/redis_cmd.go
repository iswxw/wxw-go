/*
@Time : 2022/3/22 17:40
@Author : weixiaowei
@File : redis_cmd
*/
package redis_cmd

import (
	"context"
	"framework/w01_redis/init_redis"
)

// 相关文档：https://pkg.go.dev/github.com/go-redis/redis

// 1. 根据前缀获取Key
func Keys(ctx context.Context, prefix string) ([]string, error) {
	return init_redis.RDB.Keys(ctx, "prefix*").Result()
}

// 2. 自定义命令
func Do(ctx context.Context) (interface{}, error) {
	return init_redis.RDB.Do(ctx, "set", "key", "value").Result()
}

// 3. 按照通配符删除key
func DelKey(ctx context.Context, prefix string) {
	// cursor 光标
	// count 数量统计
	iter := init_redis.RDB.Scan(ctx, 0, "prefix*", 0).Iterator()
	for iter.Next(ctx) {
		err := init_redis.RDB.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
