/*
@Time : 2022/3/28 11:05
@Author : weixiaowei
@File : demo_zset
*/
package main

import (
	"context"
	"fmt"
	"framework/w01_redis/init_redis"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisExample2()
}

func redisExample2() {

	ctx := context.Background()
	if err := init_redis.InitClient(); err != nil {
		return
	}

	zsetKey := "language_rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD
	num, err := init_redis.RDB.ZAdd(ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	fmt.Println("=>把Golang的分数加10")
	newScore, err := init_redis.RDB.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	fmt.Println("=>取分数最高的3个")
	ret, err := init_redis.RDB.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = init_redis.RDB.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
