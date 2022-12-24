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
	"src/com.wxw/project_actual/module/gin-example/common/consts"
	"strings"
)

var (
	Viper = viper.New()
	Env   = consts.Dev
)

// Setup 初始化配置项
func Setup(path string) {

	Env = GetEnvironment()
	log.Println("start environment:", Env)

	if path == "" {
		dir, _ := os.Getwd()
		path = filepath.Join(dir, "../../../conf/", Env, "app.toml")
		Viper.SetConfigFile(path)
	} else {
		Viper.SetConfigFile(path)
	}

	log.Println("current path", path)

	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf(" init conf failed: %s", err))
	}

	log.Printf("[%s_conf] setup success;\n", Env)
}

// GetEnvironment 获取当前环境
func GetEnvironment() string {
	path := "src/04_project_actual/gin-example/.deploy"
	file, err := os.Open(path)
	if err != nil {
		return consts.Dev
	}
	v, err := ioutil.ReadAll(file)
	if err != nil {
		return consts.Dev
	}
	envValue := string(v)
	switch {
	case strings.Contains(envValue, "test"):
		return consts.Test
	case strings.Contains(envValue, "prod"):
		return consts.Prod
	default:
		return consts.Dev
	}
}
