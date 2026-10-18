package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	//1.创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	//2.创建上下文
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	err := rdb.Set(
		ctx,
		"temporary_message",
		"Hello Redis",
		10*time.Second,
	).Err()

	if err != nil {
		fmt.Println("写入失败: ", err)
		return
	}

	fmt.Println("写入成功， 10s后过期")

	ttl, err := rdb.TTL(ctx, "temporary_message").Result()
	if err != nil {
		fmt.Println("查询时间过期失败: ", err)
		return
	}

	fmt.Println("剩余有效时间: ", ttl)
}
