package main

import "fmt"

func Afuntion(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	ch := make(chan int) //无缓冲的channel
	go Afuntion(ch)
	ch <- 1

	// 输出结果：
	// finish
}