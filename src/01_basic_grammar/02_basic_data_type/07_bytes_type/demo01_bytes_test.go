// @Time : 2023/6/1 16:12
// @Author : xiaoweiwei
// @File : demo01_bytes_test

package _7_bytes_type

import (
	"bytes"
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("TestHelloWorld")
	log.Println("result = ", buf.String())
}
