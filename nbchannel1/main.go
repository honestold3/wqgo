package main

import "fmt"

func Afuntion(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	ch := make(chan int) //无缓冲的channel
	//只是把这两行的代码顺序对调一下
	ch <- 1
	go Afuntion(ch)

	// 输出结果：
	// 死锁，无结果

	//代码分析：首先创建一个无缓冲的channel,　然后在主协程里面向channel　ch 中通过ch<-1命令写入数据，
	//则此时主协程阻塞，就无法执行下面的go Afuntions(ch),自然也就无法解除主协程的阻塞状态，则系统死锁
	//
	//总结：
	//对于无缓存的channel,放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中，防止死锁的发生；
	//同时应该先利用go 开一个协程对channel进行操作，此时阻塞该go 协程，然后再在主协程中进行channel的相反操作（与go 协程对channel进行相反的操作），
	//实现go 协程解锁．即必须go协程在前，解锁协程在后．
}