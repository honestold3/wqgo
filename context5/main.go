package main

import (
	"context"
	"net/http"
	"time"
	"fmt"
)

func main() {
	cx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	req = req.WithContext(cx)
	ch := make(chan error)

	go func() {
		kkk, err := http.DefaultClient.Do(req)
		fmt.Println(kkk.Status)
		select {
		case <-cx.Done():
			// Already timedout
		default:
			ch <- err
		}
	}()

	// Simulating user cancel request
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()
	select {
	case err := <-ch:
		if err != nil {
			// HTTP error
			panic(err)
		}
		fmt.Println("is ok",)
	case <-cx.Done():
		//panic(cx.Err())
		fmt.Println("error")
	}

}
