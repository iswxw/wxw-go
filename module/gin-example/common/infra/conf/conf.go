/*
@Time: 2022/12/24 15:35
@Author: wxw
@File: conf
*/
package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Dev  = "dev"
	Test = "test"
	Prod = "prod"
)

var (
	Viper = viper.New()
	Env   = Dev
)

// Setup 初始化配置项
func Setup(path string) {

	Env = GetEnvironment()
	log.Println("start environment:", Env)

	if path == "" {
		dir, _ := os.Getwd()
		path = filepath.Join(dir, "/module/gin-example/conf/", Env, "app.toml")

		// windows 本地golang编译
		if runtime.GOOS == "windows" {
			_, fn, _, _ := runtime.Caller(0)
			dir := filepath.Dir(fn)
			path = filepath.Join(dir, "../../../conf/", Env, "app.toml")
		}

		Viper.SetConfigFile(path)
	} else {
		Viper.SetConfigFile(path)
	}

	log.Println("current path", path)

	err := Viper.ReadInConfig()
	if err != nil {
		panic(any(fmt.Sprintf(" setup conf failed: %s", err)))
	}

	log.Printf("[%s_conf] setup success;\n", Env)
}

// GetEnvironment 获取当前环境
func GetEnvironment() string {
	path := "module/gin-example/.deploy"
	file, err := os.Open(path)
	if err != nil {
		return Dev
	}
	v, err := ioutil.ReadAll(file)
	if err != nil {
		return Dev
	}
	envValue := string(v)
	switch {
	case strings.Contains(envValue, "test"):
		return Test
	case strings.Contains(envValue, "prod"):
		return Prod
	default:
		return Dev
	}
}
