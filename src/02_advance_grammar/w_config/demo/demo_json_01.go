/*
@Time: 2022/1/12 23:10
@Author: wxw
@File: case
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_config/tools"
	"sync"
)

type Configs map[string]json.RawMessage

var instanceOnce sync.Once
var configPath = tools.GetFilePath(tools.ConfigFilePath, "config.json")
var CONFIG = new(Config)

// 主配置,解析json的结构体
type System struct {
	Mode string `json:"mode" ini:"mode"`
}

type Log struct {
	Prefix  string `json:"prefix" ini:"prefix"`
	LogFile bool   `json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
	Stdout  string `json:"stdout" ini:"stdout"`
	File    string `json:"file" ini:"file"`
}

// 配置结构体
type Config struct {
	System System `json:"system" ini:"system"`
	Log    Log    `json:"log" ini:"log"`
}

func main() {

	// 1. 加载文件路径
	filePath := tools.GetFilePath(tools.ConfigFilePath, "config.json")
	log.Println("filePath = ", filePath)

	// 2. 初始化文件
	Init(filePath)

	// 3. 查看配置
	configInfo, _ := json.Marshal(CONFIG)
	log.Printf("config info : \n %s \n", string(configInfo))
}

// 1.  初始化，只能运行一次
func Init(path string) {
	if CONFIG != nil && path != configPath {
		log.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
	}
	instanceOnce.Do(func() {
		mainConfig := LoadConfig(path)
		CONFIG = mainConfig
	})
}

// 2. 从配置文件中载入json字符串
func LoadConfig(path string) *Config {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}

	config := &Config{}
	if err = json.Unmarshal(buf, config); err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return config
}

// 6. 根据key获取对应的值，如果值为struct，则继续反序列化
func (cfg Configs) GetConfig(key string, config interface{}) error {
	c, ok := cfg[key]
	if ok {
		return json.Unmarshal(c, config)
	} else {
		return fmt.Errorf("fail to get cfg with key: %s", key)
	}
}
