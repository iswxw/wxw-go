/*
@Time : 2022/3/22 17:38
@Author : weixiaowei
@File : distribution_lock_normal.go
*/
package distribution_lock

import (
	"context"
	"framework/w_redis/init"
	"github.com/spf13/cast"
	"time"
)

type IDistributionLock interface {
	// 锁定资源，锁定duration时间
	LockResource(ctx context.Context, resourceId string, duration time.Duration) bool
	// 释放资源
	UnlockResource(ctx context.Context, resourceId string) bool
}

func NewDistributionLock() IDistributionLock {
	return &RedisLock{}
}

type RedisLock struct{}

func (l *RedisLock) LockResource(ctx context.Context, resourceId string, duration time.Duration) bool {
	result, err := SetRedisLock(ctx, resourceId, int(duration.Seconds()), "lock")
	if err != nil {
		return false
	}
	return result
}

func (l *RedisLock) UnlockResource(ctx context.Context, resourceId string) bool {
	result, err := DelRedisLock(ctx, resourceId)
	if err != nil {
		return false
	}
	return result
}

/**
 * 设置RedisLock
 */
func SetRedisLock(ctx context.Context, key string, expireTime int, data string) (ok bool, err error) {
	boolCmd := init.RedisClient.SetNX(ctx, key, data, cast.ToDuration(expireTime))
	if boolCmd != nil {
		return false, err
	}

	return true, nil
}

/**
 * 删除RedisLock
 */
func DelRedisLock(ctx context.Context, key string) (ok bool, err error) {
	boolCmd := init.RedisClient.Del(ctx, key)
	if boolCmd != nil {
		return false, err
	}

	return true, nil
}
