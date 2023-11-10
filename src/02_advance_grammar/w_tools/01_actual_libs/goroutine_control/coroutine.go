package sugar

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"runtime/debug"
)

const (
	coroutineNameContextKey = "coroutine_name"
	anonymousName           = "Anonymous"
)

// 异步工具
func newCoroutine(ctx context.Context, name string, fn func(ctx context.Context)) {
	spawned := context.WithValue(ctx, coroutineNameContextKey, name)
	go func() {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				err := fmt.Errorf(string(debug.Stack()))
				funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
				log.Printf("newCoroutine err = %s funName = %s\n", err, funcName)
			}
		}()
		log.Printf("newCoroutine is pawn a coroutine, coroutine_name= %s\n", name)
		fn(spawned)
	}()
}

func Go(ctx context.Context, fn func(ctx context.Context)) {
	newCoroutine(ctx, anonymousName, fn)
}
func GoNamed(ctx context.Context, name string, fn func(ctx context.Context)) {
	newCoroutine(ctx, name, fn)
}
