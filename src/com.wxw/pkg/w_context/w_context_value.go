/*
@Time: 2021/11/28 15:31
@Author: wxw
@File: w_context_value
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value :", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
