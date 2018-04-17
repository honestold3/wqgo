package main

import (
	"context"
	"fmt"
	"time"
)

func Cdd(ctx context.Context) int {
	fmt.Println(ctx.Value("NLJB"))
	fmt.Println(ctx.Value("haha"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -3
	default:
		// 没有结束 ... 执行 ...
		return 1
	}
}

func Bdd(ctx context.Context) int {
	fmt.Println(ctx.Value("HELLO"))
	fmt.Println(ctx.Value("WROLD"))
	ctx = context.WithValue(ctx, "NLJB", "NULIJIABEI")
	go fmt.Println(Cdd(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -2
	default:
		// 没有结束 ... 执行 ...
		return 1
	}
}

func Add(ctx context.Context) int {
	fmt.Println(ctx.Value("haha"))
	ctx = context.WithValue(ctx, "HELLO", "WROLD")
	ctx = context.WithValue(ctx, "WROLD", "HELLO")
	go fmt.Println(Bdd(ctx))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		return -1
	default:
		// 没有结束 ... 执行 ...
		return 1
	}
}

func main() {

	// 自动取消(定时取消)
	{
		timeout := 3 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		ctx = context.WithValue(ctx,"haha","hehe")
		fmt.Println(Add(ctx))
	}
	// 手动取消
	//  {
	//      ctx, cancel := context.WithCancel(context.Background())
	//      go func() {
	//          time.Sleep(2 * time.Second)
	//          cancel() // 在调用处主动取消
	//      }()
	//      fmt.Println(Add(ctx))
	//  }

}
