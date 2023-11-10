// @Time : 2023/4/3 15:34
// @Author : xiaoweiwei
// @File : impl

package secret_algorithm

import "context"

type AlgorithmSecret struct {
	impl ISecretAlgorithm
}

// setAlgorithmImpl 设置算法实现
func (a *AlgorithmSecret) setAlgorithmImpl(i ISecretAlgorithm) {
	a.impl = i
}

func (a *AlgorithmSecret) Encrypt(ctx context.Context, data string) (string, error) {
	return a.impl.Encrypt(ctx, data)
}
func (a *AlgorithmSecret) Decrypt(ctx context.Context, data string) (string, error) {
	return a.impl.Decrypt(ctx, data)
}

func (a *AlgorithmSecret) Sign(ctx context.Context, data string) (string, error) {
	return a.impl.Sign(ctx, data)
}
func (a *AlgorithmSecret) VerifySign(ctx context.Context, data string, sign string) (bool, error) {
	return a.impl.VerifySign(ctx, data, sign)
}
