package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var m sync.Mutex
	a := 0
	go func() {
		m.Lock()
		defer m.Unlock()
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			a = a + 1
		}

	}()

	go func() {
		m.Lock()
		for i := 0; i < 5000; i++ {
			a = a + 1
		}
		wg.Done()
		m.Unlock()
	}()

	wg.Wait()
	fmt.Println(a)
}
