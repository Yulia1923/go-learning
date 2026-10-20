package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("任务完成")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)
	defer cancel()

	err := doWork(ctx)

	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("任务超时")
		return
	}

	if errors.Is(err, context.Canceled) {
		fmt.Println("任务被取消")
		return
	}

	if err != nil {
		fmt.Println("其他错误: ", err)
		return
	}

	fmt.Println("任务执行成功")
}
