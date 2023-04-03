// @Time : 2023/4/3 15:50
// @Author : xiaoweiwei
// @File : rsa

package impl

import (
	"context"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/secret_algorithm"
)

type AlgorithmRSA struct {
	secret_algorithm.SecretAlgorithm
}

func NewAlgorithmRSA() *AlgorithmRSA {
	return &AlgorithmRSA{}
}

func (a *AlgorithmRSA) Encrypt(ctx context.Context, data string) (string, error) {

	return "", nil
}
func (a *AlgorithmRSA) Decrypt(ctx context.Context, data string) (string, error) {
	return "", nil
}

func (a *AlgorithmRSA) Sign(ctx context.Context, data string) (string, error) {
	return "", nil
}
func (a *AlgorithmRSA) VerifySign(ctx context.Context, data string, sign string) (bool, error) {
	return false, nil
}
