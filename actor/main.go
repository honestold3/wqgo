package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/goconsole"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	//props := actor.FromInstance(&HelloActor{})
	props := actor.FromProducer(func() actor.Actor { return &HelloActor{} })
	pid := actor.Spawn(props)
	pid.Tell(Hello{Who: "Roger"})
	console.ReadLine()
}