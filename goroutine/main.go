package main

import (
	"fmt"
	"time"
	"sync"
)

func task1(s string) {
	fmt.Println("task1 start")
	time.Sleep(1*time.Second)
	fmt.Println("task1 end",s)
}

func task2() {
	fmt.Println("task2 start")
	time.Sleep(2*time.Second)
	fmt.Println("task2 end")
}

func task11(wg *sync.WaitGroup) {
	fmt.Println("task1 start")
	time.Sleep(1*time.Second)
	fmt.Println("task1 end")
	wg.Done()
}

func task22(wg *sync.WaitGroup) {
	fmt.Println("task2 start")
	time.Sleep(2*time.Second)
	fmt.Println("task2 end")
	wg.Done()
}

func main() {
	start := time.Now()
	//exit := make(chan int)
	//go task1("task111")
	//go task2()


	var wg sync.WaitGroup
	wg.Add(3)
	go task11(&wg)
	go task22(&wg)
	go func(s string){
		fmt.Println("task3 start")
		time.Sleep(2*time.Second)
		fmt.Println("task3 end",s)
		//exit <- 1
		wg.Done()

	}("task333")
	end := time.Now()
	fmt.Println(end.Sub(start))

	//<- exit
	wg.Wait()
}
