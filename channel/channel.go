package main

import (
	"fmt"
	"time"
)

func main() {
	event := make(chan struct{})
	c := make(chan string)

	go func(){
		s := <- c
		fmt.Println(s)
		close(event)
	}()

	time.Sleep(2 * time.Second)
	c <- "haha"
	<-event
}