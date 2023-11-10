// @Time : 2023/4/3 15:50
// @Author : xiaoweiwei
// @File : aes

package impl

import (
	"context"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/secret_algorithm"
)

type AlgorithmAES struct {
	secret_algorithm.SecretAlgorithm
}

func NewAlgorithmAES() *AlgorithmAES {
	return &AlgorithmAES{}
}

func (a *AlgorithmAES) Encrypt(ctx context.Context, data string) (string, error) {

	return "", nil
}
func (a *AlgorithmAES) Decrypt(ctx context.Context, data string) (string, error) {
	return "", nil
}

func (a *AlgorithmAES) Sign(ctx context.Context, data string) (string, error) {
	return "", nil
}
func (a *AlgorithmAES) VerifySign(ctx context.Context, data string, sign string) (bool, error) {
	return false, nil
}
