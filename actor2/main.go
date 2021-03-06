package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/goconsole"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Started, initialize actor here")
	case *actor.Stopping:
		fmt.Println("Stopping, actor is about shut down")
	case *actor.Stopped:
		fmt.Println("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		fmt.Println("Restarting, actor is about restart")
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	props := actor.FromProducer(func() actor.Actor { return &HelloActor{} })
	pid := actor.Spawn(props)
	pid.Tell(Hello{Who: "Roger"})

	//why wait?
	//Stop is a system message and is not processed through the user message mailbox
	//thus, it will be handled _before_ any user message
	//we only do this to show the correct order of events in the console
	time.Sleep(100 * time.Microsecond)
	pid.Stop()

	console.ReadLine()

}
