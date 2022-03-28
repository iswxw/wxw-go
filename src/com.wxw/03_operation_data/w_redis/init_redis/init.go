/*
@Time : 2022/3/28 11:38
@Author : weixiaowei
@File : init_redis
*/
package init_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RDB *redis.Client
)

// 初始化连接
func InitClient() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()

	return err
}
