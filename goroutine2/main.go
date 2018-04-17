package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func main() {
	//kankan()
	exampleAtomic()
}

func kankan(){
	values := []int{1,2,3}
	for _, v:= range values {
		go func(){
			fmt.Println("kankan:",v)
		}()
	}
	time.Sleep(1*time.Second)

	fmt.Println("---------------------")

}

func kankan1(){
	values := []int{1,2,3}
	for _, v := range values{
		go func(x int) {
			fmt.Println("kankan:",x)
		}(v)
	}
	time.Sleep(1*time.Second)

	fmt.Println("---------------------")
}

func exampleAtomic() {
	var ops uint64

	for i := 0; i < 3; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops counter", opsFinal)
}


