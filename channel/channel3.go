package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int,3)
	a <- 1
	close(a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println("-------------------")

	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func(){
		defer wg.Done()
		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func(){
		defer wg.Done()
		defer close(c)
		for i:=0;i < 3;i++ {
			send <- i
		}
	}()

	wg.Wait()
}
