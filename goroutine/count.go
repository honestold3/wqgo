package main

import (
	"fmt"
	"sync"
)

func count(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	ch <- 1
	fmt.Println("Counting")
}

func kankan(i int,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	for i:=0;i<10;i++{
		go kankan(i,&wg)
	}

 	//go kankan(1,&wg)

	wg.Wait()

	//chs := make([] chan int, 10)
	//for i:=10;i<10;i++ {
	//	fmt.Println(i)
	//	chs[i] = make(chan int)
	//	go count(chs[i],&wg)
	//}
	//for _,ch := range(chs) {
	//	<-ch
	//}

}
