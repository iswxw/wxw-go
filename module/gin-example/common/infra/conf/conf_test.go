/*
@Time: 2022/12/24 16:05
@Author: wxw
@File: conf_test
*/
package conf

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestSetup(t *testing.T) {
	Setup("")
	port := Viper.GetInt("server.port")
	log.Println("port = ", port)
}

// TestPathConfig 测试文件路径
func TestPathConfig(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	newPath := filepath.Join(dir, "../../../conf/", "dev")
	log.Println(newPath)

	log.Println(GetEnvironment())
}
