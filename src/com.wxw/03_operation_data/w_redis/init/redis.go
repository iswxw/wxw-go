/*
@Time : 2022/3/22 17:47
@Author : weixiaowei
@File : redis
*/
package init

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var RedisClient *redis.Client
var ctx = context.Background()
var mutex sync.Mutex

func NewRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //ip:port
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}
