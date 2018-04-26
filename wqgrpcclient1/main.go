package main

import (
	"log"
	"time"

	"runtime"
	"sync"
	inf "wqgo/wqgrpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	wg sync.WaitGroup
)

const (
	networkType = "tcp"
	server      = "127.0.0.1"
	port        = "50051"
	parallel    = 5  //连接并行度
	times       = 1000 //每连接请求次数
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	//并行请求
	for i := 0; i < int(parallel); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exe()
		}()
	}
	wg.Wait()

	log.Printf("time taken: %.2fs", time.Now().Sub(currTime).Seconds())
}

func exe() {
	//建立连接
	conn, _ := grpc.Dial(server+":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := inf.NewGreeterClient(conn)

	for i := 0; i < int(times); i++ {
		getUser(client)
	}
}

func getUser(client inf.GreeterClient) {
	var request inf.HelloRequest
	request.Name = "hhh"

	r, err := client.SayHello(context.Background(), &request) //调用远程方法

	//判断返回结果是否正确
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
