/*
@Time : 2022/3/29 17:33
@Author : weixiaowei
@File : demo_string_join 字符串拼接
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 1.字符串拼接分析：https://geektutu.com/post/hpg-string-concat.html
// 2.综合易用性和性能，一般推荐使用 strings.Builder 来拼接字符串。
func main() {

}

// 1. 使用 加法拼接
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// 2. 使用 fmt.Sprintf
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// 3. 使用 strings.Builder
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 4. 使用 bytes.Buffer
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

// 5. 使用 []byte
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
