package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Sum []int

type Result struct {
	value int
	count int
	time  time.Time
}

type Message struct {
	v1   int
	v2   int
	v3   int
	time time.Time
}

type SumActor struct{}

func (state *SumActor) Receive(context actor.Context) {
	r := 10000
	//masterActor := context.Spawn(actor.FromProducer(func() actor.Actor { return &MasterActor{} }))
	switch msg := context.Message().(type) {
	case Message:
		//fmt.Printf("Hello %v,%v\n", msg.v1, msg.v2)
		start := (r/msg.v2)*(msg.v1-1) + 1
		end := (r / msg.v2) * msg.v1
		flag := msg.v1

		c := Calculate(start, end, flag)
		//fmt.Println("c::", c)
		context.Parent().Tell(Result{value: c, count: 1, time: msg.time})

		//masterActor.Tell(Result{value: c,count :1,time: msg.time})
	}
}

func doWork(context actor.Context) {
	r := 10000
	switch msg := context.Message().(type) {
	case Message:
		start := (r/msg.v2)*(msg.v1-1) + 1
		end := (r / msg.v2) * msg.v1
		flag := msg.v1

		c := Calculate(start, end, flag)
		//fmt.Println("c::", c)
		context.Parent().Tell(Result{value: c, count: 1, time: msg.time})
	}
}

type Start struct {
	start string
	time  time.Time
}

type MasterActor struct {
	count     int
	sum       int
	starttime time.Time
}

func (state *MasterActor) Receive(context actor.Context) {
	state.starttime = time.Now()
	ncpu := 4
	printActor := context.Spawn(actor.FromProducer(func() actor.Actor { return &PrintActor{} }))
	//PrintActor := actor.Spawn(router.NewRoundRobinPool(5).WithFunc(act))

	sumActor := context.Spawn(actor.FromProducer(func() actor.Actor { return &SumActor{} }))
	//sumActor := context.Spawn(router.NewRoundRobinPool(ncpu).WithProducer(func() actor.Actor { return &SumActor{} }))
	//kankan1 := context.Spawn(router.NewRoundRobinPool(ncpu).WithFunc(doWork))

	switch msg := context.Message().(type) {

	case Start:
		if msg.start == "start" {
			fmt.Println(msg.start, ncpu)
			for i := 1; i < ncpu+1; i++ {
				sumActor.Tell(Message{v1: i, v2: ncpu, v3: 1, time: msg.time})
				//kankan1.Tell(Message{v1: i, v2: ncpu, v3: 1, time: msg.time})
			}

		}
	case Result:
		state.count++
		state.sum += msg.value
		if state.count == ncpu {
			printActor.Tell(Print{sum: strconv.Itoa(state.sum), startTime: msg.time})
		}
	}

}

type PrintActor struct {
}

type Print struct {
	sum       string
	startTime time.Time
}

func (state *PrintActor) Receive(context actor.Context) {

	switch msg := context.Message().(type) {
	case Print:
		fmt.Println("总数为：", msg.sum, "；所花时间为：", (time.Now().Sub(msg.startTime)), "。")
	}
}

func Calculate(start, end int, flag int) int {
	cal := 0
	for i := start; i <= end; i++ {
		for j := 1; j <= 3000000; j++ {
		}
		cal += i
	}
	fmt.Println("flag :", flag, ".")
	return cal
}

func main() {
	props := actor.FromProducer(func() actor.Actor { return &MasterActor{} })
	pid := actor.Spawn(props)
	pid.Tell(Start{start: "start", time: time.Now()})
	pid.Stop()
	console.ReadLine()
}
