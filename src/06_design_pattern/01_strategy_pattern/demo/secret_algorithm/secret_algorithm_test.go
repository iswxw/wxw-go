// @Time : 2023/4/3 15:38
// @Author : xiaoweiwei
// @File : impl_test

package secret_algorithm

import (
	"context"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/secret_algorithm/impl"
	"testing"
)

func TestEncrypt(t *testing.T) {
	algorithm := AlgorithmSecret{}
	algorithm.setAlgorithmImpl(impl.NewAlgorithmAES())
	ctx := context.Background()

	encrypt, err := algorithm.Encrypt(ctx, "明文数据")
	println("err = ", err)
	println("encrypt = ", encrypt)
}
