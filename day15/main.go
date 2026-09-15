package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex

	var wg sync.WaitGroup //创建一个 WaitGroup（等待组）

	wg.Add(2) //要等待 2 个 Goroutine 完成

	go func() {
		defer wg.Done()

		mu.Lock()
		count++
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		mu.Lock()
		count++
		mu.Lock()
	}()

	wg.Wait() //等到 WaitGroup 的计数器变成 0，再继续往下执行
	fmt.Println(count)
}
