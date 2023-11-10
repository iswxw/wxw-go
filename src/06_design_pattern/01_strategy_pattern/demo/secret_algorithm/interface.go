// @Time : 2023/4/3 15:34
// @Author : xiaoweiwei
// @File : interface

package secret_algorithm

import "context"

type ISecretAlgorithm interface {
	Encrypt(ctx context.Context, data string) (string, error)
	Decrypt(ctx context.Context, data string) (string, error)
	Sign(ctx context.Context, data string) (string, error)
	VerifySign(ctx context.Context, data string, sign string) (bool, error)
}
