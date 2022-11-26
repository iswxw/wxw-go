/*
@Time : 2022/3/22 17:55
@Author : weixiaowei
@File : distribution_lock_mutex.go
*/
package distribution_lock

import (
	"context"
	"framework/w01_redis/init_redis"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex
var ctx = context.Background()

// 加锁
func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	bool, err := init_redis.RDB.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}

// 释放锁
func UnLock(key string) int64 {
	nums, err := init_redis.RDB.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
