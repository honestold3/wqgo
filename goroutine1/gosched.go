package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan int)

	go func(){
		defer func(){
			close(exit) //和 exit<-1 功能类似
			fmt.Println("exit")
		}()

		go func() {
			fmt.Println("b")
		}()

		fmt.Println("a1")
		runtime.Gosched()
		fmt.Println("a2")
	}()
	<-exit
}

