package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	a,b := make(chan int),make(chan int)

	wg.Add(2)
	go func(){
		defer wg.Done()
		for{
			var name string
			var x int
			var ok bool
			select {
				case x,ok = <-a:
					name = "a"
				case x,ok = <-b:
					name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name,x)
		}
	}()

	go func(){
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i:=0;i<10;i++{
			select {
				case a <- i:
				case b <- i*10:
			}
		}
	}()
	wg.Wait()

	fmt.Println("-------------")

	t := time.NewTicker(1 * time.Second)
	for {
		select {
			case <- t.C:
				fmt.Println("aa")
		}
	}
}