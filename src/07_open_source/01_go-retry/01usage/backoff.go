// @Time : 2023/2/17 17:30
// @Author : xiaoweiwei
// @File : backoff

package main

import (
	"context"
	"github.com/sethvargo/go-retry"
	"time"
)

func main() {
	ctx := context.Background()

	b := retry.NewFibonacci(1 * time.Second)
	b = retry.WithMaxDuration(5*time.Second, b)

	if err := retry.Do(ctx, b, func(_ context.Context) error {
		// TODO: logic here
		return nil
	}); err != nil {
		// handle error
	}
}
