package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

//WithTimeout 等价于 WithDeadline(parent, time.Now().Add(timeout))
func timeoutHandler() {
	 ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	 go doTimeOutStuff(ctx)
	//go doStuff(ctx)

	time.Sleep(10 * time.Second)

	//cancel()

}

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok { //设置了deadl
			logg.Printf("deadline set")
			if time.Now().After(deadline) {
				logg.Printf(ctx.Err().Error())
				return
			}

		}

		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("end")
}

