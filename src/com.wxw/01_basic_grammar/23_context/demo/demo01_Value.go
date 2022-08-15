// @Time : 2022/8/15 15:36
// @Author : xiaoweiwei
// @File : demo01_Value

package main

import (
	"context"
	"fmt"
)

// 应用场景：设置上下文的信息
func main() {
	ctx := context.Background()
	process(ctx)

	// 重新设置上下文的值
	ctx = context.WithValue(ctx, "traceId", "qcrao-2019")
	process(ctx)
}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
