package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // 多次调用只执行一次
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// 输出结果：
// Only once
