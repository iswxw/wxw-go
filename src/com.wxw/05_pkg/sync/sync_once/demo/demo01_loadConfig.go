/*
@Time: 2021/12/5 22:06
@Author: wxw
@File: demo
@link：https://geektutu.com/post/hpg-sync-once.html
*/
package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}

// 读取配置
func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_ROOT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Printf("init_redis config: %+v", config)
	})
	return config
}
