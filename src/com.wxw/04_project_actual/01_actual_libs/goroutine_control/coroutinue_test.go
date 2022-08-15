package sugar

import (
	"context"
	"fmt"
	"testing"
)

// TestGo 测试捕获 goroutinue 异常
func TestGo(t *testing.T) {
	var testFunc = func(ctx context.Context) {
		fmt.Println("aaaa")
		panic("123")
	}

	//testFunc(context.Background())
	Go(context.Background(), testFunc)
}

func TestGoNamed(t *testing.T) {
	var testFunc = func(ctx context.Context) {
		fmt.Println("aaaa")
		panic("123")
	}

	//testFunc(context.Background())
	GoNamed(context.Background(), "testFunc", testFunc)
}
