package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Sum []int

func (s Sum) Calculate(count, start, end int, flag string, ch chan int) {
	cal := 0

	for i := start; i <= end; i++ {
		for j := 1; j <= 3000000; j++ {
		}
		cal += i
	}

	s[count] = cal
	fmt.Println("flag :", flag, ".")
	ch <- count
}

func (s Sum) LetsGo() {

	//ncpu := runtime.NumCPU()
	ncpu := 8
	const RANGE = 10000
	var ch = make(chan int,100)
	//fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(ncpu)
	for i := 0; i < ncpu; i++ {
		go s.Calculate(i, (RANGE/ncpu)*i+1, (RANGE/ncpu)*(i+1), strconv.Itoa(i+1), ch)
	}

	for i := 0; i < ncpu; i++ {
		<-ch
	}
}

func main() {
	var s Sum = make([]int, 20, 20)
	var sum int = 0
	var startTime = time.Now()

	s.LetsGo()

	for _, v := range s {
		sum += v
	}

	fmt.Println("总数为：", sum, "；所花时间为：", (time.Now().Sub(startTime)), "秒。")
}
