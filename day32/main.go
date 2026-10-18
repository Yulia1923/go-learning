package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

const nullCacheValue = "__NULL__"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ctx := context.Background()

	// 1. 连接 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	// 2. 连接 MySQL
	db, err := sql.Open(
		"mysql",
		"root:USuk1923@tcp(127.0.0.1:3306)/go_mysql_demo",
	)
	if err != nil {
		log.Fatal("初始化 MySQL 失败:", err)
	}
	defer db.Close()

	// 3. 检查 MySQL 是否可以连接
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("连接 MySQL 失败:", err)
	}

	key := "user:1001"

	// 4. 先查询 Redis
	value, err := rdb.Get(ctx, key).Result()

	if err == nil {
		var user User

		err := json.Unmarshal([]byte(value), &user)
		if err != nil {
			log.Println("解析缓存JSON失败:", err)
			return
		}

		fmt.Printf(
			"缓存命中: id=%d, name=%s, age=%d\n",
			user.ID, user.Name, user.Age,
		)
		return
	}

	// 5. 只有 redis.Nil 才代表缓存未命中
	if err != redis.Nil {
		log.Println("查询Redis失败:", err)
		return
	}

	fmt.Println("缓存未命中，开始查询 MySQL")

	// 6. 缓存未命中，回源 MySQL
	var id int
	var name string
	var age int

	err = db.QueryRowContext(
		ctx,
		"SELECT id, name, age FROM users WHERE id = ?",
		1001,
	).Scan(&id, &name, &age)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("MySQL 中也没有这个用户")

			err := rdb.Set(
				ctx,
				key,
				nullCacheValue,
				time.Minute,
			).Err()

			if err != nil {
				log.Println("写入空值缓存失败:", err)
			}

			return
		}
		log.Fatal("查询 MySQL 失败:", err)
	}

	user := User{
		ID:   id,
		Name: name,
		Age:  age,
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Println("转换JSON失败:", err)
		return
	}

	err = rdb.Set(ctx, key, data, 5*time.Minute).Err()
	if err != nil {
		log.Println("写入Redis缓存失败:", err)
	} else {
		fmt.Println("MySQL查询结果已写入Redis缓存")
	}

	fmt.Printf(
		"MySQL 查询成功: id=%d, name=%s, age=%d\n",
		id, name, age,
	)
}
