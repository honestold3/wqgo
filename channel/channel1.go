package main

import (
	"fmt"
)

func main() {
	c := make(chan int,3)

	c <- 1
	c <- 2
	fmt.Println(<-c)
	c <- 3
	c <- 4
	fmt.Println(<-c)

	a := make(chan int, 1)
	b := make(chan int, 3)
	a <- 1
	b <- 11
	b <- 2
	fmt.Println("a:", len(a), cap(a))
	fmt.Println("b:", len(b), cap(b))

	value, ok := <-b
	fmt.Println(value, ok)
	value1, ok1 := <-b
	fmt.Println(value1, ok1)
}