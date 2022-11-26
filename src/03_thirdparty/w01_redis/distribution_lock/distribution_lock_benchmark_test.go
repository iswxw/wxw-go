/*
@Time : 2022/3/28 14:05
@Author : weixiaowei
@File : distribution_lock_benchmark
*/
package distribution_lock

import (
	"fmt"
	"framework/w01_redis/init_redis"
	"testing"
)

func BenchmarkLock(b *testing.B) {

	if err := init_redis.InitClient(); err != nil {
		return
	}

	// 加锁
	flag := "set"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println("BenchMark Test:", Lock(flag))
	}

}
