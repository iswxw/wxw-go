/*
@Time: 2022/12/21 18:11
@Author: wxw
@File: main_test
*/
package runtime

import (
	"runtime"
	"testing"
)

func TestTestHelloWorld(t *testing.T) {
	println(runtime.NumCPU())
}
